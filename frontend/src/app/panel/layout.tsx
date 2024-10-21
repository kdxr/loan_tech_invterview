import LayoutProvider from "@/providers/LayoutProvider";
import React from "react";

type Props = {} & React.PropsWithChildren;

export default function layout({ children }: Props) {
	return <LayoutProvider>{children}</LayoutProvider>;
}
