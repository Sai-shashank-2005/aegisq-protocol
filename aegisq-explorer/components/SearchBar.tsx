"use client"

import {useState} from "react"
import {useRouter} from "next/navigation"

export default function SearchBar(){

const [query,setQuery]=useState("")
const router=useRouter()

function search(){

if(!query) return

if(/^\d+$/.test(query)){
router.push(`/block/${query}`)
}else{
router.push(`/txhash/${query}`)
}

}

return(

<div className="flex gap-2">

<input
className="bg-gray-900 border border-gray-700 px-3 py-2 rounded w-64"
placeholder="Search block height or tx hash"
value={query}
onChange={e=>setQuery(e.target.value)}
/>

<button
onClick={search}
className="bg-blue-600 px-4 py-2 rounded"
>
Search
</button>

</div>
)
}