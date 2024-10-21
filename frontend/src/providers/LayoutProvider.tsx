"use client";
import React, { useCallback, useMemo, useState } from "react";
import { Card, Layout, Menu } from "antd";
import { usePathname, useRouter } from "next/navigation";
import MenuItems from "@/components/Menu/MenuItems";
import type { MenuProps } from "antd";

const { Header, Content, Sider } = Layout;

type Props = {} & React.PropsWithChildren;

export default function LayoutProvider({ children }: Props) {
	const router = useRouter();
	const pathname = usePathname();

	const [collapsed, setCollapsed] = useState<boolean>(false);

	const currentMenu = useMemo((): React.Key => {
		if (!pathname) return "";

		const menu = MenuItems.find((item) => item.path === pathname);

		if (menu && menu.key) return menu.key;

		return "";
	}, [pathname]);

	const handleChangePage = useCallback<Required<MenuProps>["onClick"]>(
		(info) => {
			const menu = MenuItems.find((item) => item.key === info.key);

			if (!menu || !menu.path) return;

			router.push(menu.path);
		},
		[router, MenuItems]
	);

	return (
		<main className="w-full min-h-screen flex gap-3 bg-background">
			<Layout style={{ minHeight: "100vh" }}>
				<Sider
					collapsible
					collapsed={collapsed}
					onCollapse={(value) => setCollapsed(value)}
				>
					<div className="demo-logo-vertical" />
					<Menu
						theme="dark"
						selectedKeys={[currentMenu as string]}
						mode="inline"
						items={MenuItems}
						onClick={handleChangePage}
					/>
				</Sider>
				<Layout>
					<Content style={{ padding: 10 }}>
						<Card classNames={{ body: "h-full" }}>{children}</Card>
					</Content>
				</Layout>
			</Layout>
		</main>
	);
}
