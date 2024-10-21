"use client";
import { DeleteOutlined, EditOutlined, EyeOutlined } from "@ant-design/icons";
import {
	Button,
	Col,
	DatePicker,
	Flex,
	Form,
	Input,
	InputNumber,
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
import dayjs, { Dayjs } from "dayjs";
import { TCustomer } from "../customers/page";
import ModalInformation from "@/components/Loans/ModalInformation";

type Props = {};

type TLoan = {
	customerEmail: string;
	customerId: number;
	customerName: string;
	customerTel: string;
	interestRate: number;
	loanAmount: number;
	loanEndDate: Date | Dayjs;
	loanStartDate: Date | Dayjs;
	loanId: number;
};

type ModalStateKey = "" | "create" | "edit";

export default function page({}: Props) {
	const [isModalOpen, setIsModalOpen] = useState(false);
	const [modalStateKey, setModalStateKey] = useState<ModalStateKey>("");
	const [loans, setLoans] = useState<TLoan[]>();
	const [customers, setCustomers] = useState<TCustomer[]>();

	const [isModalInformationOpen, setIsModalInformationOpen] = useState(false);
	const [loanInformationId, setLoanInformationId] = useState<number>();

	const [search, setSearch] = useState<string>();

	const [form] = Form.useForm<TLoan>();

	const loanFilters = useMemo<TLoan[]>(() => {
		if (!Array.isArray(loans)) return [];

		if (!search) return loans;

		return loans.filter(
			(customer) =>
				customer.customerName.includes(search) ||
				customer.customerTel.includes(search) ||
				customer.customerEmail.includes(search)
		);
	}, [loans, search]);

	const customerOptions = useMemo(() => {
		if (!Array.isArray(customers)) return [];

		return customers.map((customer) => ({
			label: customer.name,
			value: customer.id,
		}));
	}, [customers]);

	const columns: TableProps<TLoan>["columns"] = [
		{
			title: "Customer Name",
			dataIndex: "customerName",
			key: "customerName",
		},
		{
			title: "Customer Email",
			dataIndex: "customerEmail",
			key: "customerEmail",
		},
		{
			title: "Customer Telephone Number",
			dataIndex: "customerTel",
			key: "customerTel",
		},
		{
			title: "Interest Rate",
			dataIndex: "interestRate",
			key: "interestRate",
		},
		{
			title: "Loan Amount",
			dataIndex: "loanAmount",
			key: "loanAmount",
		},
		{
			title: "Start Date",
			key: "loanStartDate",
			dataIndex: "loanStartDate",
			render: (value) => {
				return dayjs(value).format("DD-MM-YYYY");
			},
		},
		{
			title: "End Date",
			key: "loanEndDate",
			dataIndex: "loanEndDate",
			render: (value) => {
				return dayjs(value).format("DD-MM-YYYY");
			},
		},
		{
			title: "Action",
			key: "action",
			align: "center",
			render: (_, record) => (
				<Space size="middle">
					<Button
						color="primary"
						variant="text"
						icon={<EyeOutlined />}
						onClick={() => handleOPenModalInformation(record.loanId)}
					/>
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
						onClick={() => handleDelete(record.loanId)}
					/>
				</Space>
			),
		},
	];

	const handleCreate = () => {
		setIsModalOpen(true);
		setModalStateKey("create");
	};

	const handleEdit = (values: TLoan) => {
		setIsModalOpen(true);
		setModalStateKey("edit");

		form.setFieldsValue({
			...values,
			loanStartDate: dayjs(values.loanStartDate),
			loanEndDate: dayjs(values.loanEndDate),
		});
	};

	const handleCancel = () => {
		setIsModalOpen(false);
		form.resetFields();
	};

	const handleFinish = async (values: TLoan) => {
		try {
			const model: {
				loanAmount: number;
				interestRate: number;
				startDate: Date;
				endDate: Date;
				customerId?: number;
				id?: number;
			} = {
				loanAmount: values.loanAmount,
				interestRate: values.interestRate,
				startDate: dayjs(values.loanStartDate).toDate(),
				endDate: dayjs(values.loanEndDate).toDate(),
			};

			if (modalStateKey === "create") {
				model.customerId = values.customerId;
			} else if (modalStateKey === "edit") {
				model.id = values.loanId;
			}

			const url: string =
				modalStateKey === "create"
					? "http://localhost/manage-loans/create"
					: "http://localhost/manage-loans/update";

			const response = await axios.post(url, model);

			message.success(response.data.message);

			await handleFetchloans();

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
				`http://localhost/manage-loans/delete/${id}`
			);

			message.success(response.data.message);

			await handleFetchloans();
		} catch (error: AxiosError | any) {
			message.error(error?.response?.data?.message ?? "Failed");
		}
	};

	const handleFetchloans = async () => {
		try {
			const { data } = await axios.get("http://localhost/manage-loans/lists");

			setLoans(data.data);
		} catch (error) {
			setLoans([]);
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

	const handleOPenModalInformation = (id: number) => {
		setIsModalInformationOpen(true);
		setLoanInformationId(id);
	};

	const handleCloseInformation = () => {
		setIsModalInformationOpen(false);
		setLoanInformationId(undefined);
	};

	useEffect(() => {
		handleFetchloans();
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
			<Table<TLoan>
				rowKey="loanId"
				columns={columns}
				dataSource={loanFilters}
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
					<Form.Item name="loanId" hidden>
						<Input hidden />
					</Form.Item>

					<Form.Item
						label="Customer "
						name="customerId"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<Select
							options={customerOptions}
							disabled={modalStateKey === "edit"}
						/>
					</Form.Item>

					<Form.Item
						label="Loan Amount"
						name="loanAmount"
						rules={[
							{
								required: true,
								validator: (_, value) => {
									return value > 0 ? Promise.resolve() : Promise.reject();
								},
								message: "Interest rate can't less than 0",
							},
						]}
						required
					>
						<InputNumber style={{ width: "100%" }} />
					</Form.Item>

					<Form.Item
						label="Interest Rate"
						name="interestRate"
						rules={[
							{
								required: true,
								validator: (_, value) => {
									return value > 0 && value <= 100
										? Promise.resolve()
										: Promise.reject();
								},
								message: "Interest rate can't less than 0 or more than 100",
							},
						]}
						required
					>
						<InputNumber style={{ width: "100%" }} />
					</Form.Item>

					<Form.Item
						label="Start Date"
						name="loanStartDate"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<DatePicker style={{ width: "100%" }} />
					</Form.Item>

					<Form.Item
						label="End Date"
						name="loanEndDate"
						rules={[
							{
								required: true,
							},
						]}
						required
					>
						<DatePicker style={{ width: "100%" }} />
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
			<Modal
				open={isModalInformationOpen}
				footer={null}
				mask
				maskClosable
				onCancel={handleCloseInformation}
				centered
				destroyOnClose
			>
				<ModalInformation loanId={loanInformationId} />
			</Modal>
		</Flex>
	);
}
