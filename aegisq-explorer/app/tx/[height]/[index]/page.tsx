"use client"

import { useParams } from "next/navigation"
import { useEffect, useState } from "react"
import { getTx } from "@/lib/api"
import Loader from "@/components/Loader"

export default function TransactionPage() {

  const params = useParams()
  const height = Number(params.height)
  const index = Number(params.index)

  const [tx, setTx] = useState<any>(null)

  useEffect(() => {
    getTx(height, index).then(setTx)
  }, [height, index])

  if (!tx) {
    return <Loader label="Decoding Transaction Payload" />
  }

  return (
    <div className="relative w-full px-8 py-10 space-y-12">

      {/* 🔥 GRID BACKGROUND */}
      <div className="absolute inset-0 opacity-[0.03] pointer-events-none
        bg-[linear-gradient(#fff_1px,transparent_1px),linear-gradient(90deg,#fff_1px,transparent_1px)]
        bg-[size:40px_40px]" />

      {/* 🔥 HEADER */}
      <div className="relative z-10 flex items-center justify-between">

        <div>
          <h1 className="text-3xl font-semibold text-white">
            Transaction #{index}
          </h1>

          <p className="text-gray-500 text-sm mt-1">
            Block {height} • Cryptographic Inspection Layer
          </p>
        </div>

        {/* STATUS BADGE */}
        <div className="px-4 py-1.5 rounded-full text-xs
          border border-green-500/30 text-green-400
          bg-green-500/10 backdrop-blur-md
          shadow-[0_0_12px_rgba(34,197,94,0.2)]">
          VERIFIED
        </div>

      </div>

      {/* 🔥 MAIN CARD */}
      <div className="relative z-10 w-full
        bg-gradient-to-br from-[#0b1220] to-[#0a0f1a]
        border border-gray-800 rounded-2xl
        p-10 space-y-8
        shadow-[0_0_60px_rgba(0,0,0,0.8)]">

        {/* SENDER */}
        <Row label="Sender">
          <span className="text-blue-400 font-medium">
            {tx.sender_id}
          </span>
        </Row>

        {/* ALGORITHM */}
        <Row label="Algorithm">
          <span className="text-gray-300">
            {tx.algorithm}
          </span>
        </Row>

        {/* 🔥 HASH */}
        <div className="border-t border-gray-800 pt-6">

          <p className="text-xs text-gray-500 mb-3 uppercase tracking-wide">
            Data Hash
          </p>

          <div className="flex items-center justify-between gap-4
            bg-black/40 border border-gray-800 rounded-xl px-4 py-3">

            <span className="font-mono text-xs text-blue-400 break-all">
              {tx.data_hash}
            </span>

            <button
              onClick={() => navigator.clipboard.writeText(tx.data_hash)}
              className="text-xs px-3 py-1 rounded-md border border-gray-700
                hover:border-blue-500 hover:text-blue-400 transition">
              Copy
            </button>

          </div>

        </div>

        {/* METADATA */}
        <Row label="Metadata">
          <span className="text-gray-300">{tx.metadata}</span>
        </Row>

        {/* TIMESTAMP */}
        <Row label="Timestamp">
          <span className="text-gray-400">
            {new Date(tx.timestamp * 1000).toLocaleString()}
          </span>
        </Row>

      </div>

    </div>
  )
}

/* 🔹 ROW COMPONENT */
function Row({ label, children }: any) {
  return (
    <div className="flex justify-between items-center border-t border-gray-800 pt-6">

      <span className="text-xs text-gray-500 uppercase tracking-wide">
        {label}
      </span>

      <div className="text-sm text-right max-w-[60%] break-words">
        {children}
      </div>

    </div>
  )
}