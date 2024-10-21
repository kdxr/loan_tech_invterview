import type { Config } from "tailwindcss";

const config: Config = {
	content: [
		"./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
		"./src/components/**/*.{js,ts,jsx,tsx,mdx}",
		"./src/app/**/*.{js,ts,jsx,tsx,mdx}",
		"./src/providers/**/*.{js,ts,jsx,tsx,mdx}",
		// "./src/**/*.{js,ts,jsx,tsx,mdx}",
	],
	theme: {
		screens: {
			xs: "480px",
			sm: "576px",
			md: "768px",
			lg: "992px",
			xl: "1200px",
			xxl: "1600px",
		},
		extend: {
			fontFamily: {
				kanit: ["var(--font-kanit)", "sans-serif"],
			},
		},
	},
	darkMode: "class",
};
export default config;
