import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/sidebar";
import Header from "@/components/header";
import HeaderMobile from "@/components/header-mobile";
import axios from "axios";
import { Toaster } from "@/components/ui/toaster";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="h-full">
      <body className={`${inter.className} flex overflow-y-hidden h-full`}>
        <Sidebar />
        <main className="flex flex-1">
          <div className="flex w-full flex-col overflow-y-hidden">
            <Header />
            <HeaderMobile />
            <div className="md:p-10 p-4 overflow-auto">{children}</div>
            <Toaster />
          </div>
        </main>
      </body>
    </html>
  );
}
