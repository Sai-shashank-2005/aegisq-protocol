"use client"

import {useEffect,useState} from "react"
import {getBlocks} from "@/lib/api"
import BlockCard from "@/components/BlockCard"

export default function Blocks(){

const [blocks,setBlocks]=useState<any[]>([])

useEffect(()=>{
getBlocks().then(setBlocks)
},[])

return(

<div>

<h1 className="text-3xl mb-6">Blocks</h1>

<div className="grid grid-cols-2 gap-4">

{blocks.map(b=>(
<BlockCard key={b.height} block={b}/>
))}

</div>

</div>

)
}