"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { Search } from "lucide-react"

export default function SearchBar() {

  const [query, setQuery] = useState("")
  const router = useRouter()

  function search() {
    if (!query.trim()) return

    if (/^\d+$/.test(query)) {
      router.push(`/block/${query}`)
    } else {
      router.push(`/txhash/${query}`)
    }
  }

  return (

    <div className="relative w-full max-w-md">

      {/* 🔥 INPUT */}
      <input
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        onKeyDown={(e) => e.key === "Enter" && search()}
        placeholder="Search block height or tx hash..."
        className="
          w-full pl-10 pr-24 py-2.5
          bg-gray-900/80 backdrop-blur
          border border-gray-800
          rounded-xl text-sm text-white
          placeholder:text-gray-500
          focus:outline-none focus:border-blue-500/50
          transition
        "
      />

      {/* 🔍 ICON */}
      <Search
        size={16}
        className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500"
      />

      {/* 🔥 BUTTON (embedded) */}
      <button
        onClick={search}
        className="
          absolute right-1 top-1/2 -translate-y-1/2
          px-3 py-1.5 rounded-lg text-xs font-medium
          bg-blue-500/10 border border-blue-500/30
          text-blue-400 hover:bg-blue-500/20
          transition
        "
      >
        Search
      </button>

    </div>

  )
}