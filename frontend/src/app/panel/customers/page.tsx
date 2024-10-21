"use client";
import { DeleteOutlined, EditOutlined } from "@ant-design/icons";
import {
	Button,
	Col,
	Flex,
	Form,
	Input,
	message,
	Modal,
	Row,
	Select,
	Space,
	Switch,
	Table,
	TableProps,
	Tag,
} from "antd";
import React, { useEffect, useMemo, useState } from "react";
import axios, { AxiosError } from "axios";

type Props = {};

export interface TCustomer {
	id: number;
	key: number;
	name: string;
	email: string;
	telphoneNumber: string;
	statusActive: boolean;
	createdAt: Date;
	updatedAt: Date;
	sex: string;
	address: string;
	tumbon: string;
	district: string;
	province: string;
}

type ModalStateKey = "" | "create" | "edit";

export default function page({}: Props) {
	const [isModalOpen, setIsModalOpen] = useState(false);
	const [modalStateKey, setModalStateKey] = useState<ModalStateKey>("");
	const [customers, setCustomers] = useState<TCustomer[]>();

	const [search, setSearch] = useState<string>();

	const [form] = Form.useForm<TCustomer>();

	const customerFilters = useMemo<TCustomer[]>(() => {
		if (!Array.isArray(customers)) return [];

		if (!search) return customers;

		return customers.filter(
			(customer) =>
				customer.name.includes(search) ||
				customer.telphoneNumber.includes(search) ||
				customer.email.includes(search)
		);
	}, [customers, search]);

	const columns: TableProps<TCustomer>["columns"] = [
		{
			title: "Name",
			dataIndex: "name",
			key: "name",
		},
		{
			title: "Email",
			dataIndex: "email",
			key: "email",
		},
		{
			title: "Telephone Number",
			dataIndex: "telphoneNumber",
			key: "telphoneNumber",
		},
		{
			title: "Status",
			key: "statusActive",
			dataIndex: "statusActive",
			render: (value: boolean) => (
				<>
					<Tag color={value ? "green" : "red"}>
						{value ? "Active" : "Inactive"}
					</Tag>
				</>
			),
		},
		{
			title: "Action",
			key: "action",
			align: "center",
			render: (_, record) => (
				<Space size="middle">
					<Button
						color="default"
						variant="text"
						icon={<EditOutlined />}
						onClick={() => handleEdit(record)}
					/>
					<Button
						color="danger"
						variant="text"
						icon={<DeleteOutlined />}
						onClick={() => handleDelete(record.id)}
					/>
				</Space>
			),
		},
	];

	const handleCreate = () => {
		setIsModalOpen(true);
		setModalStateKey("create");
	};

	const handleEdit = (values: TCustomer) => {
		setIsModalOpen(true);
		setModalStateKey("edit");
		form.setFieldsValue(values);
	};

	const handleCancel = () => {
		setIsModalOpen(false);
		form.resetFields();
	};

	const handleFinish = async (values: TCustomer) => {
		try {
			const model = {
				id: values.id ?? undefined,
				email: values.email,
				telphoneNumber: values.telphoneNumber,
				name: values.name,
				sex: values.sex,
				address: values.address,
				tumbon: values.tumbon,
				district: values.district,
				province: values.province,
				statusActive: values.statusActive,
			};

			const url: string =
				modalStateKey === "create"
					? "http://localhost/manage-customers/create"
					: "http://localhost/manage-customers/update";

			const response = await axios.post(url, model);

			message.success(response.data.message);

			await handleFetchCustomers();

			setIsModalOpen(false);
			form.resetFields();
		} catch (error: AxiosError | any) {
			message.error(error?.response?.data?.message ?? "Failed");
			console.log(error);
		} finally {
		}
	};

	const handleDelete = async (id: number) => {
		try {
			const response = await axios.delete(
				`http://localhost/manage-customers/delete/${id}`
			);

			message.success(response.data.message);

			await handleFetchCustomers();
		} catch (error: AxiosError | any) {
			message.error(error?.response?.data?.message ?? "Failed");
		}
	};

	const handleFetchCustomers = async () => {
		try {
			const { data } = await axios.get(
				"http://localhost/manage-customers/lists"
			);

			setCustomers(data.data);
		} catch (error) {
			setCustomers([]);
		}
	};

	const handleSearch: React.ChangeEventHandler<HTMLInputElement> = (e) => {
		setSearch(e.target.value);
	};

	useEffect(() => {
		handleFetchCustomers();
	}, []);

	return (
		<Flex gap={10} vertical>
			<Row gutter={[16, 16]}>
				<Col xs={{ span: 20 }}>
					<Input placeholder="Search..." onChange={handleSearch} />
				</Col>
				<Col xs={{ span: 4 }}>
					<Button color="primary" block onClick={handleCreate}>
						Add
					</Button>
				</Col>
			</Row>
			<Table<TCustomer>
				rowKey="id"
				columns={columns}
				dataSource={customerFilters}
			/>
			<Modal
				open={isModalOpen}
				footer={null}
				mask
				maskClosable
				onCancel={handleCancel}
				centered
			>
				<Form form={form} layout="vertical" onFinish={handleFinish}>
					<Form.Item name="id" hidden>
						<Input hidden />
					</Form.Item>
					<Form.Item
						label="Name"
						name="name"
						rules={[
							{
								required: true,
								transform: (val) => val?.trim(),
								message: "Please input name",
							},
						]}
						required
					>
						<Input />
					</Form.Item>
					<Form.Item
						label="Email"
						name="email"
						rules={[
							{
								required: true,
								type: "email",
								message: "Invalid email address",
							},
						]}
						required
					>
						<Input type="email" />
					</Form.Item>
					<Form.Item
						label="Telephone Number"
						name="telphoneNumber"
						rules={[
							{
								required: true,
								len: 10,
							},
							{
								pattern: new RegExp(/^[0-9]+$/),
								message: "Invalid telephone number",
							},
						]}
						required
					>
						<Input style={{ width: "100%" }} />
					</Form.Item>
					<Form.Item
						label="Gender"
						name="sex"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<Select
							options={[
								{
									value: "male",
									label: "Male",
								},
								{
									value: "female",
									label: "Female",
								},
							]}
						/>
					</Form.Item>
					<Form.Item
						label="Address"
						name="address"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<Input.TextArea />
					</Form.Item>
					<Form.Item
						label="Tumbon"
						name="tumbon"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<Input />
					</Form.Item>
					<Form.Item
						label="District"
						name="district"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<Input />
					</Form.Item>
					<Form.Item
						label="Province"
						name="province"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<Input />
					</Form.Item>
					<Form.Item
						label="Status"
						name="statusActive"
						valuePropName="checked"
						initialValue={true}
						required
					>
						<Switch
							checkedChildren="Active"
							unCheckedChildren="Inactive"
							className="flex"
						/>
					</Form.Item>

					<Row justify="end" gutter={16}>
						<Col>
							<Button type="default" onClick={handleCancel}>
								Cancel
							</Button>
						</Col>

						<Col>
							<Form.Item>
								<Button type="primary" htmlType="submit">
									Confirm
								</Button>
							</Form.Item>
						</Col>
					</Row>
				</Form>
			</Modal>
		</Flex>
	);
}
