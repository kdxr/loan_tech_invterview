import {
	Button,
	Form,
	InputNumber,
	message,
	Table,
	TableProps,
	Tabs,
	TabsProps,
} from "antd";
import axios from "axios";
import dayjs from "dayjs";
import React, { useEffect, useState } from "react";

type Props = {
	loanId: number | undefined;
};

interface FormValues {
	payAmount: number;
	principleAmount: number;
	interestAmount: number;
}

interface TPayment {
	createdAt: Date;
	payAmount: number;
	principleAmount: number;
	interestAmount: number;
}

const TabPaymentList = ({ loanId }: Props) => {
	const [paymentList, setPaymentList] = useState<TPayment[]>([]);

	const columns: TableProps<TPayment>["columns"] = [
		{
			title: "No",
			dataIndex: "no",
			render: (_, __, index) => index + 1,
		},
		{
			title: "Pay amount",
			dataIndex: "payAmount",
			key: "payAmount",
		},
		{
			title: "Principle amount",
			dataIndex: "principleAmount",
			key: "principleAmount",
		},
		{
			title: "Interest Amount",
			dataIndex: "interestAmount",
			key: "interestAmount",
		},

		{
			title: "Date",
			key: "createdAt",
			dataIndex: "createdAt",
			render: (value) => {
				return dayjs(value).format("DD-MM-YYYY HH:mm:ss");
			},
		},
	];

	const handleFetchLoanInformation = async () => {
		try {
			const response = await axios.get(
				`http://localhost/manage-loans/info/${loanId}`
			);

			setPaymentList(response.data.data.payments);
		} catch (error) {
			message.error("Failed");
		}
	};

	useEffect(() => {
		if (loanId) {
			handleFetchLoanInformation();
		}
	}, [loanId]);

	return (
		<Table<TPayment>
			rowKey={(record, index) => String(index)}
			columns={columns}
			dataSource={paymentList}
		/>
	);
};

const TabCreate = ({ loanId }: { loanId: number | undefined }) => {
	const [form] = Form.useForm<FormValues>();

	const payAmount = Form.useWatch("payAmount", form);
	const principleAmount = Form.useWatch("principleAmount", form);
	const interestAmount = Form.useWatch("interestAmount", form);

	const handleCreatePayment = async (values: FormValues) => {
		try {
			const model = {
				loanId: loanId,
				payAmount: values.payAmount,
				principleAmount: values.principleAmount,
				interestAmount: values.interestAmount,
			};

			const response = await axios.post(
				"http://localhost/manage-payments/create",
				model
			);

			form.resetFields();

			message.success(response.data.message);
		} catch (error: any) {
			message.error(error.response.data.message);
		}
	};

	return (
		<Form form={form} onFinish={handleCreatePayment} layout="vertical">
			<Form.Item
				name="payAmount"
				label="Pay Amount"
				rules={[
					{
						required: true,
						validator: (_, value) => {
							return value > 0 ? Promise.resolve() : Promise.reject();
						},
						message: "Interest rate can't less than 0",
					},
				]}
			>
				<InputNumber style={{ width: "100%" }} />
			</Form.Item>
			<Form.Item
				name="principleAmount"
				label="Principle Amount"
				rules={[
					{
						required: true,
						validator: (_, value) => {
							return value > 0 ? Promise.resolve() : Promise.reject();
						},
						message: "Interest rate can't less than 0",
					},
					{
						validator: (_, value) => {
							if (payAmount && principleAmount + interestAmount > payAmount) {
								return Promise.reject();
							}

							return Promise.resolve();
						},
						message: "Principle amount can't greater than pay amount",
					},
				]}
				dependencies={["payAmount", "interestAmount"]}
			>
				<InputNumber style={{ width: "100%" }} />
			</Form.Item>
			<Form.Item
				name="interestAmount"
				label="Interest Amount"
				rules={[
					{
						required: true,
						validator: (_, value) => {
							return value > 0 ? Promise.resolve() : Promise.reject();
						},
						message: "Interest rate can't less than 0",
					},
					{
						validator: (_, value) => {
							if (payAmount && principleAmount + interestAmount > payAmount) {
								return Promise.reject();
							}

							return Promise.resolve();
						},
						message: "Interest amount can't greater than pay amount",
					},
				]}
				dependencies={["payAmount", "principleAmount"]}
			>
				<InputNumber style={{ width: "100%" }} />
			</Form.Item>
			<Form.Item>
				<Button type="primary" htmlType="submit" style={{ width: "100%" }}>
					Create
				</Button>
			</Form.Item>
		</Form>
	);
};

export default function ModalInformation({ loanId }: Props) {
	const tabItems: TabsProps["items"] = [
		{
			key: "lists",
			label: "List Payment",
			children: <TabPaymentList loanId={loanId} />,
		},
		{
			key: "create",
			label: "Create Payment",
			children: <TabCreate loanId={loanId} />,
		},
	];

	return (
		<div>
			<Tabs items={tabItems} destroyInactiveTabPane />
		</div>
	);
}
