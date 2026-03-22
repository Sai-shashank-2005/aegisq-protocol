"use client"

import { useEffect, useState } from "react"
import { useParams, useRouter } from "next/navigation"
import { getBlock } from "@/lib/api"
import Loader from "@/components/Loader"

export default function BlockPage() {
  const params = useParams()
  const router = useRouter()
  const height = Number(params.height)

  const [block, setBlock] = useState<any>(null)

  useEffect(() => {
    getBlock(height).then(setBlock)
  }, [height])

  if (!block) {
    return <Loader label={`Syncing Block #${height}`} />
  }

  return (
    <div className="w-full px-8 py-6 space-y-10 relative">

      {/* 🔥 GRID BACKGROUND */}
      <div className="absolute inset-0 -z-10 opacity-[0.05] 
        bg-[linear-gradient(#00ffcc_1px,transparent_1px),linear-gradient(90deg,#00ffcc_1px,transparent_1px)] 
        bg-[size:40px_40px]" 
      />

      {/* HEADER */}
      <div>
        <h1 className="text-3xl font-semibold tracking-tight">
          Block #{block.height}
        </h1>
        <p className="text-gray-500 text-sm mt-1">
          Distributed consensus snapshot
        </p>
      </div>

      {/* 🚀 STATUS PANEL */}
      <div className="relative rounded-xl border border-green-500/30 bg-gradient-to-r from-green-500/10 to-transparent p-6 overflow-hidden">

        <div className="absolute inset-0 bg-green-500/20 blur-3xl opacity-20" />

        <div className="relative z-10 flex justify-between items-center">

          {/* LEFT */}
          <div>
            <p className="text-green-400 font-semibold text-lg flex items-center gap-2">
              <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse" />
              FINALIZED BLOCK
            </p>

            <p className="text-sm text-gray-400 mt-1">
              Leader: {block.leader} • Quorum{" "}
              {block.consensus.quorum.received}/
              {block.consensus.quorum.required}
            </p>
          </div>

          {/* RIGHT */}
          <div className="flex items-center gap-3">

            <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse" />

            <span className="text-green-400 text-sm font-medium tracking-wide">
              FINALIZED
            </span>

            <span className="text-[10px] px-2 py-1 rounded-md bg-green-500/10 border border-green-500/20 text-green-300 tracking-wider">
              BFT
            </span>

          </div>

        </div>

        <div className="absolute bottom-0 left-0 w-full h-[1px] bg-gradient-to-r from-green-500/0 via-green-400/40 to-green-500/0" />

      </div>

      {/* 📊 METRICS */}
      <div className="grid grid-cols-3 gap-6">

        <Metric 
          title="Leader" 
          value={block.leader} 
          highlight 
        />

        <Metric
          title="Quorum"
          value={`${block.consensus.quorum.received}/${block.consensus.quorum.required}`}
        />

        <Metric
          title="Consensus"
          value={block.consensus.status}
          success
        />

      </div>

      {/* 📦 TRANSACTIONS */}
      <div className="bg-gradient-to-b from-gray-900 to-gray-950 border border-gray-800 rounded-xl overflow-hidden">

        <div className="p-6 border-b border-gray-800 flex justify-between items-center">

          <h2 className="text-lg font-semibold tracking-wide">
            Transactions
          </h2>

          <div className="text-sm text-gray-500">
            {block.transactions.length} total
          </div>

        </div>

        <table className="w-full text-sm">

          <thead className="text-gray-400 border-b border-gray-800">
            <tr>
              <th className="p-4 text-left">Index</th>
              <th className="p-4 text-left">Sender</th>
              <th className="p-4 text-right">Action</th>
            </tr>
          </thead>

          <tbody>

            {block.transactions.slice(0, 50).map((tx: any, i: number) => (

              <tr
                key={i}
                onClick={() => router.push(`/tx/${block.height}/${i}`)}
                className="group border-b border-gray-800 hover:bg-blue-500/5 cursor-pointer transition duration-200"
              >

                <td className="p-4 text-blue-400 font-medium">
                  #{i}
                </td>

                <td className="p-4 text-gray-300 group-hover:text-white transition">
                  {tx.sender_id}
                </td>

                <td className="p-4 text-right text-gray-500 group-hover:text-blue-400 transition">
                  View →
                </td>

              </tr>

            ))}

          </tbody>

        </table>

      </div>

    </div>
  )
}

function Metric({ title, value, highlight = false, success = false }: any) {
  return (
    <div className="relative bg-gradient-to-b from-gray-900 to-gray-950 border border-gray-800 rounded-xl p-6 hover:border-blue-500/30 transition overflow-hidden">

      <div className="absolute inset-0 opacity-0 hover:opacity-20 bg-blue-500 blur-2xl transition" />

      <p className="text-gray-500 text-xs uppercase tracking-wider">
        {title}
      </p>

      <p
        className={`text-xl mt-3 font-semibold tracking-wide ${
          highlight
            ? "text-purple-400"
            : success
            ? "text-green-400"
            : "text-white"
        }`}
      >
        {value}
      </p>
    </div>
  )
}