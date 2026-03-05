"use client"

import Link from "next/link"
import SearchBar from "./SearchBar"

export default function Navbar(){

return(

<div className="w-full border-b border-gray-800 bg-black">

<div className="max-w-6xl mx-auto flex items-center justify-between py-4">

<div className="flex gap-6 items-center">

<Link href="/" className="text-xl font-bold">
AegisQ Explorer
</Link>

<Link href="/blocks" className="text-gray-400 hover:text-white">
Blocks
</Link>

</div>

<SearchBar/>

</div>
</div>
)
}