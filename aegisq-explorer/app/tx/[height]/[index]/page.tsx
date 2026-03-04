"use client"

import {useParams} from "next/navigation"
import {useEffect,useState} from "react"
import {getTx} from "../../../../lib/api"

export default function TxPage(){

  const params = useParams()

  const height = Number(params.height)
  const index = Number(params.index)

  const [tx,setTx] = useState<any>(null)

  useEffect(()=>{
    getTx(height,index).then(setTx)
  },[height,index])

  if(!tx) return <div>Loading...</div>

  return(

    <div>

      <h1 className="text-3xl mb-6">
        Transaction
      </h1>

      <div className="bg-gray-900 p-6 rounded space-y-2">

        <div>Sender: {tx.sender_id}</div>
        <div>Algorithm: {tx.algorithm}</div>
        <div>Data Hash: {tx.data_hash}</div>
        <div>Metadata: {tx.metadata}</div>
        <div>Timestamp: {tx.timestamp}</div>

      </div>

    </div>
  )
}