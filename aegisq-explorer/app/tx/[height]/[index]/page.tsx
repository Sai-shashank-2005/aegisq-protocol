"use client"

import { useParams } from "next/navigation"
import { useEffect, useState } from "react"
import { getTransaction } from "../../../../lib/api"

export default function TxPage() {

  const params = useParams()

  const height = Number(params.height)
  const index = Number(params.index)

  const [tx, setTx] = useState<any>(null)

  useEffect(() => {
    async function load() {
      const data = await getTransaction(height, index)
      setTx(data)
    }

    load()

  }, [height, index])

  if (!tx) return <div>Loading transaction...</div>

  return (
    <div>

      <h1 className="text-3xl mb-6">
        Transaction {index}
      </h1>

      <div className="bg-gray-900 border border-gray-800 p-6 rounded">

        <p><b>Sender:</b> {tx.sender_id}</p>
        <p><b>Algorithm:</b> {tx.algorithm}</p>
        <p><b>Metadata:</b> {tx.metadata}</p>
        <p><b>Timestamp:</b> {tx.timestamp}</p>

      </div>

    </div>
  )
}