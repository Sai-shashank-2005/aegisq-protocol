"use client"

import {useEffect,useState} from "react"
import {getStatus,getBlocks} from "@/lib/api"
import Hero from "@/components/Hero"
import MetricCard from "@/components/MetricCard"

export default function Dashboard(){

const [status,setStatus]=useState<any>()

useEffect(()=>{
getStatus().then(setStatus)
},[])

if(!status) return <div>Loading...</div>

return(

<div className="space-y-8">

<Hero
title="AegisQ Network Dashboard"
subtitle="Real-time hybrid BFT blockchain monitoring"
/>

{/* METRICS */}

<div className="grid grid-cols-4 gap-6">

<MetricCard
title="Latest Block"
value={status.height}
color="text-white"
/>

<MetricCard
title="Validators"
value="4"
color="text-green-400"
/>

<MetricCard
title="Block Size"
value="10000 TX"
color="text-blue-400"
/>

<MetricCard
title="TPS"
value="1200"
color="text-purple-400"
/>

</div>

</div>

)

}