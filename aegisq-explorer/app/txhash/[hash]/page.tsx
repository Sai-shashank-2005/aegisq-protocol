"use client"

import { useParams } from "next/navigation"
import { useEffect,useState } from "react"
import { getTxHash } from "@/lib/api"

export default function TxHash(){

const params = useParams()
const hash = String(params.hash)

const [tx,setTx] = useState<any>()

useEffect(()=>{
getTxHash(hash).then(setTx)
},[hash])

if(!tx) return <div>Loading...</div>

return(

<div>

<h1 className="text-3xl mb-6">Transaction</h1>

<div className="bg-gray-900 p-6 rounded space-y-3">

<div>Sender: {tx.sender_id}</div>
<div>Algorithm: {tx.algorithm}</div>
<div>Data Hash: {tx.data_hash}</div>
<div>Metadata: {tx.metadata}</div>
<div>Timestamp: {tx.timestamp}</div>

</div>

</div>

)

}