import type { MenuProps } from "antd";
import { FileOutlined, UserOutlined } from "@ant-design/icons";

type MenuItem = Required<MenuProps>["items"][number] & { path: string };

const MenuItems: MenuItem[] = [
	{
		key: "customers",
		path: "/panel/customers",
		icon: <UserOutlined />,
		label: "Customers",
	},
	{
		key: "loans",
		path: "/panel/loans",
		icon: <FileOutlined />,
		label: "Loans",
	},
];

export default MenuItems;
