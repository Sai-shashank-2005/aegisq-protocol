"use client"

import {useEffect,useState} from "react"
import {getStatus,getBlocks} from "@/lib/api"

export default function Dashboard(){

const [status,setStatus]=useState<any>()
const [blocks,setBlocks]=useState<any[]>([])

useEffect(()=>{

getStatus().then(setStatus)
getBlocks().then(setBlocks)

},[])

if(!status) return <div>Loading...</div>

return(

<div className="space-y-8">

<h1 className="text-3xl font-bold">
Network Overview
</h1>

{/* ================= METRICS ================= */}

<div className="grid grid-cols-4 gap-6">

<Metric title="Latest Block" value={status.height}/>

<Metric title="Validators" value="4"/>

<Metric title="Block Size" value="10000"/>

<Metric title="TPS" value="1200"/>

</div>

{/* ================= RECENT BLOCKS ================= */}

<div className="bg-gray-900 border border-gray-800 rounded-xl">

<div className="p-4 border-b border-gray-800 font-semibold">
Recent Blocks
</div>

<table className="w-full text-left">

<thead className="text-gray-400 border-b border-gray-800">
<tr>
<th className="p-3">Height</th>
<th className="p-3">Transactions</th>
<th className="p-3">Hash</th>
</tr>
</thead>

<tbody>

{blocks.slice(0,10).map((b)=>(
<tr
key={b.height}
className="border-b border-gray-800 hover:bg-gray-800 cursor-pointer"
onClick={()=>window.location.href=`/block/${b.height}`}
>

<td className="p-3 text-blue-400">{b.height}</td>
<td className="p-3">{b.txs}</td>
<td className="p-3 text-gray-500">
{b.hash.slice(0,20)}...
</td>

</tr>
))}

</tbody>

</table>

</div>

</div>

)

}

function Metric({title,value}:any){

return(

<div className="bg-gray-900 border border-gray-800 rounded-xl p-6">

<p className="text-gray-400 text-sm">
{title}
</p>

<p className="text-2xl mt-2 font-semibold">
{value}
</p>

</div>

)

}