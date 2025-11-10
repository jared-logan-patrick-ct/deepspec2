import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({
  weight: ['300', '400', '500', '600', '700'],
  subsets: ["latin"],
  variable: '--font-inter',
});

export const metadata = {
  title: "DeepSpec Visualizer",
  description: "Code visualization from sourcemap",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${inter.variable} font-[var(--font-inter)]`}>
        {children}
      </body>
    </html>
  );
}
