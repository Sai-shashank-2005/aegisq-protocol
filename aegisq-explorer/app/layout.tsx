import "./globals.css"
import Link from "next/link"
import SearchBar from "../components/SearchBar"

export default function RootLayout({
  children,
}:{
  children:React.ReactNode
}){

  return (
    <html>
      <body className="bg-black text-white">

        <nav className="border-b border-gray-800 p-4 flex justify-between">

          <div className="flex gap-6 text-lg">

            <Link href="/">
              AegisQ Explorer
            </Link>

            <Link href="/blocks">
              Blocks
            </Link>

          </div>

          <SearchBar/>

        </nav>

        <main className="max-w-6xl mx-auto p-6">
          {children}
        </main>

      </body>
    </html>
  )
}