"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { Search } from "lucide-react"

export default function Header() {

  const [query, setQuery] = useState("")
  const router = useRouter()

  const handleSearch = () => {
    if (!query) return

    if (!isNaN(Number(query))) {
      router.push(`/block/${query}`)
    } else {
      router.push(`/txhash/${query}`)
    }
  }

  return (
    <div className="h-16 flex items-center justify-between px-8 border-b border-gray-800 bg-black/60 backdrop-blur-lg">

      {/* LEFT */}
      <div>
        <h1 className="text-sm text-gray-500">
          AegisQ Explorer
        </h1>
      </div>

      {/* SEARCH */}
      <div className="relative w-[480px] group">

        <div className="flex items-center bg-gradient-to-r from-gray-900 to-gray-950 border border-gray-800 rounded-xl px-4 py-2 transition-all duration-200 group-focus-within:border-blue-500 group-focus-within:shadow-[0_0_20px_rgba(59,130,246,0.25)]">

          <Search size={16} className="text-gray-500 mr-3" />

          <input
            type="text"
            placeholder="Search blocks or transactions..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleSearch()}
            className="bg-transparent outline-none text-sm w-full text-white placeholder-gray-500"
          />

          <span className="text-xs text-gray-600 mr-3 hidden md:block">
            ↵
          </span>

          <button
            onClick={handleSearch}
            className="px-3 py-1.5 bg-blue-600 hover:bg-blue-500 text-white text-xs rounded-lg transition"
          >
            Search
          </button>

        </div>

      </div>

    </div>
  )
}