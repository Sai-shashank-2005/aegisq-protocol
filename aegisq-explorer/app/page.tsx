"use client"

import {useEffect,useState} from "react"
import {getStatus,getBlocks} from "@/lib/api"
import StatCard from "@/components/StatCard"
import BlockCard from "@/components/BlockCard"

export default function Home(){

const [status,setStatus]=useState<any>()
const [blocks,setBlocks]=useState<any[]>([])

useEffect(()=>{

getStatus().then(setStatus)
getBlocks().then(setBlocks)

},[])

return(

<div>

<h1 className="text-3xl mb-6">AegisQ Network</h1>

<div className="grid grid-cols-3 gap-4 mb-8">

<StatCard title="Latest Height" value={status?.height}/>
<StatCard title="Status" value={status?.status}/>
<StatCard title="Blocks Indexed" value={blocks.length}/>

</div>

<h2 className="text-xl mb-4">Recent Blocks</h2>

<div className="grid grid-cols-2 gap-4">

{blocks.slice(0,8).map(b=>
<BlockCard key={b.height} block={b}/>
)}

</div>

</div>

)
}