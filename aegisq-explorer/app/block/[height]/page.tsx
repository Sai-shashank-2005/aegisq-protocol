"use client"

import { useEffect, useState } from "react"
import { useParams } from "next/navigation"
import { getBlock } from "@/lib/api"

export default function BlockPage(){

const params = useParams()
const height = Number(params.height)

const [block,setBlock] = useState<any>()
const [page,setPage] = useState(1)

const TX_PER_PAGE = 50

useEffect(()=>{
getBlock(height).then(setBlock)
},[height])

if(!block) return <div className="p-10">Loading block...</div>

const start = (page-1)*TX_PER_PAGE
const end = start + TX_PER_PAGE

const txs = block.Transactions.slice(start,end)

const totalPages = Math.ceil(block.Transactions.length / TX_PER_PAGE)

return(

<div className="max-w-6xl mx-auto px-6">

{/* Title */}

<h1 className="text-3xl font-semibold mb-8">
Block {height}
</h1>

{/* Stats */}

<div className="grid grid-cols-3 gap-4 mb-8">

<div className="bg-gray-900 border border-gray-800 p-5 rounded">
<div className="text-gray-400 text-sm">Height</div>
<div className="text-xl">{block.Index}</div>
</div>

<div className="bg-gray-900 border border-gray-800 p-5 rounded">
<div className="text-gray-400 text-sm">Transactions</div>
<div className="text-xl">{block.Transactions.length}</div>
</div>

<div className="bg-gray-900 border border-gray-800 p-5 rounded">
<div className="text-gray-400 text-sm">View</div>
<div className="text-xl">{block.View}</div>
</div>

</div>

{/* Transactions */}

<div className="bg-gray-900 border border-gray-800 rounded">

<div className="p-5 border-b border-gray-800 flex justify-between">

<h2 className="text-lg font-semibold">Transactions</h2>

<div className="text-gray-400 text-sm">
Showing {start} - {end} of {block.Transactions.length}
</div>

</div>

<table className="w-full text-left">

<thead className="text-gray-400 border-b border-gray-800">
<tr>
<th className="p-3">Index</th>
<th className="p-3">Sender</th>
</tr>
</thead>

<tbody>

{txs.map((tx:any,i:number)=>{

const index = start + i

return(

<tr
key={index}
onClick={()=>window.location.href=`/tx/${height}/${index}`}
className="border-b border-gray-800 hover:bg-gray-800 cursor-pointer transition"
>

<td className="p-3 text-blue-400">{index}</td>

<td className="p-3">{tx.sender_id}</td>

</tr>

)

})}

</tbody>

</table>

{/* Pagination */}

<div className="flex justify-center gap-2 p-5">

<button
onClick={()=>setPage(page-1)}
disabled={page===1}
className="px-3 py-1 bg-gray-800 rounded disabled:opacity-40"
>
Prev
</button>

<div className="px-4 py-1">
Page {page} / {totalPages}
</div>

<button
onClick={()=>setPage(page+1)}
disabled={page===totalPages}
className="px-3 py-1 bg-gray-800 rounded disabled:opacity-40"
>
Next
</button>

</div>

</div>

</div>

)

}