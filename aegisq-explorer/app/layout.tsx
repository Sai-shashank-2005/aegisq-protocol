import "./globals.css"
import Link from "next/link"
import type { ReactNode } from "react"

export const metadata = {
  title: "AegisQ Explorer",
  description: "Post-Quantum Blockchain Explorer",
}

export default function RootLayout({
  children,
}: {
  children: ReactNode
}) {
  return (
    <html lang="en">
      <body className="bg-black text-white">

        <nav className="border-b border-gray-800 p-4 flex gap-6">
          <Link href="/">Dashboard</Link>
          <Link href="/blocks">Blocks</Link>
        </nav>

        <main className="max-w-6xl mx-auto p-6">
          {children}
        </main>

      </body>
    </html>
  )
}