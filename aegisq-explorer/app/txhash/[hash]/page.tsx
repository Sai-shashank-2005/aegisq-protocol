"use client"

import {useParams} from "next/navigation"
import {useEffect,useState} from "react"
import {getTxHash} from "../../../lib/api"

export default function TxHashPage(){

  const params = useParams()
  const hash = params.hash as string

  const [tx,setTx] = useState<any>(null)

  useEffect(()=>{
    getTxHash(hash).then(setTx)
  },[hash])

  if(!tx) return <div>Loading...</div>

  return(

    <div>

      <h1 className="text-3xl mb-6">
        Transaction
      </h1>

      <div className="space-y-2">

        <div>Sender: {tx.sender_id}</div>

        <div>Algorithm: {tx.algorithm}</div>

        <div>Data Hash: {tx.data_hash}</div>

        <div>Metadata: {tx.metadata}</div>

      </div>

    </div>
  )
}