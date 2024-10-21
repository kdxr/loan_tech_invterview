import type { Metadata } from "next";
import { Kanit } from "next/font/google";
import "./global.scss";
import { AntdRegistry } from "@ant-design/nextjs-registry";

const kanit = Kanit({
	weight: ["100", "200", "300", "400", "500", "600", "700", "800", "900"],
	variable: "--font-kanit",
	subsets: ["thai", "latin"],
});

export const metadata: Metadata = {
	title: "Create Next App",
	description: "Generated by create next app",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="en">
			<body className={`${kanit.className} dark antialiased`}>
				<AntdRegistry>{children}</AntdRegistry>
			</body>
		</html>
	);
}