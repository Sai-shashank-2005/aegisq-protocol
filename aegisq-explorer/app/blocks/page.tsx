"use client"

import {useEffect,useState} from "react"
import {getBlocks} from "@/lib/api"
import Hero from "@/components/Hero"

export default function BlocksPage(){

const [blocks,setBlocks] = useState<any[]>([])

useEffect(()=>{

async function load(){
const data = await getBlocks()
setBlocks(data)
}

load()

},[])

return(

<div className="space-y-8">

<Hero
title="Blocks Explorer"
subtitle="Browse all blocks produced by the AegisQ network"
/>

{blocks.length === 0 ? (

<div className="text-gray-500">
No blocks found.
</div>

) : (

<div className="grid grid-cols-3 gap-6">

{blocks.map((b)=>(

<div
key={b.height}
onClick={()=>window.location.href=`/block/${b.height}`}
className="bg-gray-900 border border-gray-800 rounded-xl p-6 hover:bg-gray-800 cursor-pointer"
>

<h2 className="text-lg font-semibold">
Block {b.height}
</h2>

<p className="text-gray-400 mt-2">
Transactions: {b.txs}
</p>

<p className="text-gray-500 text-sm mt-2">
{b.hash.slice(0,24)}...
</p>

</div>

))}

</div>

)}

</div>

)

}