"use client"

import { useEffect, useState } from "react"
import { getBlocks } from "@/lib/api"
import Loader from "@/components/Loader"
import { useRouter } from "next/navigation"

export default function BlocksPage() {

  const [blocks, setBlocks] = useState<any[]>([])
  const router = useRouter()

  useEffect(() => {
    getBlocks().then(setBlocks)
  }, [])

  if (!blocks.length) {
    return <Loader label="Syncing Blocks From Network" />
  }

  return (
    <div className="relative w-full px-8 py-10 space-y-12">

      {/* 🔥 GRID BACKGROUND */}
      <div className="absolute inset-0 opacity-[0.03] pointer-events-none
        bg-[linear-gradient(#fff_1px,transparent_1px),linear-gradient(90deg,#fff_1px,transparent_1px)]
        bg-[size:40px_40px]" />

      {/* 🔥 HEADER */}
      <div className="relative z-10">
        <h1 className="text-3xl font-semibold text-white">
          Blocks Explorer
        </h1>

        <p className="text-gray-500 text-sm mt-1">
          Real-time block production across AegisQ network
        </p>
      </div>

      {/* 🔥 BLOCK GRID */}
      <div className="relative z-10 grid grid-cols-3 gap-6">

        {blocks.map((b, i) => {

          const isLatest = i === 0

          return (
            <div
              key={b.height}
              onClick={() => router.push(`/block/${b.height}`)}
              className={`group cursor-pointer rounded-2xl p-6 transition-all duration-300
              border backdrop-blur-xl
              ${isLatest
                ? "bg-gradient-to-br from-blue-900/40 to-black border-blue-500/30 shadow-[0_0_30px_rgba(59,130,246,0.2)]"
                : "bg-gradient-to-br from-[#0b1220] to-[#0a0f1a] border-gray-800 hover:border-blue-500/40"
              } hover:scale-[1.02]`}
            >

              {/* HEADER */}
              <div className="flex justify-between items-center mb-4">

                <h2 className="text-lg font-semibold text-white">
                  Block #{b.height}
                </h2>

                {isLatest && (
                  <span className="text-xs px-2 py-1 rounded border border-green-500/30 text-green-400 bg-green-500/10">
                    Latest
                  </span>
                )}

              </div>

              {/* DATA */}
              <div className="space-y-3 text-sm">

                <div className="flex justify-between text-gray-400">
                  <span>Transactions</span>
                  <span className="text-white">{b.txs ?? 0}</span>
                </div>

                <div className="flex justify-between text-gray-400">
                  <span>Hash</span>
                  <span className="text-gray-500 truncate max-w-[120px]">
                    {b.hash}
                  </span>
                </div>

              </div>

              {/* 🔥 HOVER LINE */}
              <div className="mt-6 h-[2px] bg-gray-800 relative overflow-hidden rounded">

  {/* FULL WIDTH BASE */}
  <div className="absolute inset-0 bg-gray-800" />

  {/* SCANNING LINE */}
  <div className="absolute top-0 left-0 h-full w-full
    bg-gradient-to-r from-transparent via-blue-500 to-transparent
    opacity-0 group-hover:opacity-100
    animate-scan" />

</div>
            </div>
          )
        })}

      </div>

    </div>
  )
}

/* 🔥 FLOW ANIMATION */
<style jsx global>{`
@keyframes flow {
  0% { transform: translateX(-100%) }
  100% { transform: translateX(300%) }
}
.animate-flow {
  animation: flow 1.2s linear infinite;
}
`}</style>