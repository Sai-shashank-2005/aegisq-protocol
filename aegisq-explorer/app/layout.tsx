import "./globals.css"
import Link from "next/link"

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html>
      <body className="bg-black text-white">

        <nav className="p-4 border-b border-gray-800 flex gap-6">
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