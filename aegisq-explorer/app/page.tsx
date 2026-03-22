"use client"

import { useEffect, useState } from "react"
import { getConsensus, getLiveness } from "@/lib/api"
import Loader from "@/components/Loader"
import { Crown, Users, Activity, Zap } from "lucide-react"

export default function Dashboard() {
  const [consensus, setConsensus] = useState<any>(null)
  const [liveness, setLiveness] = useState<any>(null)

  useEffect(() => {
    const fetchData = async () => {
      const [c, l] = await Promise.all([
        getConsensus(),
        getLiveness()
      ])
      setConsensus(c)
      setLiveness(l)
    }

    fetchData()
    const i = setInterval(fetchData, 3000)
    return () => clearInterval(i)
  }, [])

  if (!consensus || !liveness) {
    return <Loader label="Syncing Network State" />
  }

  const leader = consensus.leader

  return (
    <div className="relative space-y-10 px-8 py-6">

      {/* GRID */}
      <div className="absolute inset-0 opacity-[0.03] pointer-events-none
        bg-[linear-gradient(#fff_1px,transparent_1px),linear-gradient(90deg,#fff_1px,transparent_1px)]
        bg-[size:40px_40px]" />

      {/* 🔥 HERO */}
      <div className="relative z-10 rounded-xl border border-green-500/30 bg-green-500/5 p-6 flex justify-between items-center">

        <div>
          <h1 className="text-2xl font-semibold text-white">
            AegisQ Network
          </h1>

          <p className="text-sm text-gray-400 mt-1">
            Hybrid BFT • Live Consensus Observability
          </p>

          <p className="text-green-400 text-sm mt-3">
            ● SYSTEM OPERATIONAL
          </p>
        </div>

        <div className="text-right text-sm text-gray-400">
          Leader<br />
          <span className="text-white font-medium">
            {leader}
          </span>
        </div>

      </div>

      {/* 🔥 PRIMARY METRICS */}
      <div className="grid grid-cols-3 gap-6">

        <Metric title="Leader" value={leader} icon={<Crown size={16} />} highlight />
        <Metric title="Consensus" value={consensus.status} icon={<Activity size={16} />} success />
        <Metric title="Quorum" value={`${consensus.quorum.received}/${consensus.quorum.required}`} icon={<Users size={16} />} />

      </div>

      {/* 🔥 SECONDARY METRICS */}
      <div className="grid grid-cols-3 gap-6">

        <Strip title="Latency" value="120 ms" />
        <Strip title="Leader Stability" value="HIGH" highlight />
        <Strip title="Participation" value="100%" success />

      </div>

      {/* 🔥 CONSENSUS FLOW */}
      <div className="rounded-xl border border-gray-800 bg-gray-900/60 backdrop-blur p-6">

        <p className="text-xs text-gray-400 mb-6">
          Consensus State Machine
        </p>

        <div className="flex items-center justify-between">

          <Node label="Leader" color="bg-purple-500" />
          <Flow />
          <Node label="Proposal" color="bg-blue-500" />
          <Flow />
          <Node label="Validators" color="bg-yellow-500" />
          <Flow />
          <Node label="Commit" color="bg-green-500" />

        </div>

      </div>

      {/* 🔥 VALIDATORS */}
      <div className="rounded-xl border border-gray-800 bg-gray-900/60 backdrop-blur p-6">

        <p className="text-xs text-gray-400 mb-4">
          Validators
        </p>

        <div className="grid grid-cols-4 gap-4">

          {consensus.validators.map((v: string, i: number) => {
            const isLeader = v === leader

            return (
              <div
                key={i}
                className={`
                  p-4 rounded-xl border transition-all
                  ${isLeader
                    ? "border-green-500/30 bg-green-500/5"
                    : "border-gray-800 bg-black/40"
                  }
                  hover:border-blue-500/40 hover:scale-[1.02]
                `}
              >

                <div className="flex justify-between">
                  <span className="text-sm text-white">{v}</span>
                  {isLeader && (
                    <Crown size={14} className="text-yellow-400" />
                  )}
                </div>

                <div className="text-xs text-gray-500 mt-2 flex items-center gap-2">
                  <span className="w-2 h-2 bg-green-500 rounded-full"></span>
                  Active
                </div>

              </div>
            )
          })}

        </div>

      </div>

    </div>
  )
}

/* ===== COMPONENTS ===== */

function Metric({ title, value, icon, highlight, success }: any) {
  return (
    <div className="rounded-xl border border-gray-800 bg-gray-900/60 backdrop-blur p-5 hover:border-blue-500/40 transition">

      <div className="flex items-center gap-2 text-gray-400 text-sm">
        {icon}
        {title}
      </div>

      <p className={`text-lg mt-2 font-semibold ${
        highlight ? "text-purple-400" :
        success ? "text-green-400" :
        "text-white"
      }`}>
        {value}
      </p>
    </div>
  )
}

function Strip({ title, value, highlight, success }: any) {
  return (
    <div className="rounded-xl border border-gray-800 bg-gray-900/60 backdrop-blur p-4">

      <p className="text-xs text-gray-400">{title}</p>

      <p className={`text-md mt-1 font-semibold ${
        highlight ? "text-blue-400" :
        success ? "text-green-400" :
        "text-white"
      }`}>
        {value}
      </p>
    </div>
  )
}

function Node({ label, color }: any) {
  return (
    <div className="flex flex-col items-center gap-2">
      <div className={`w-3 h-3 rounded-full ${color}`} />
      <span className="text-xs text-gray-400">{label}</span>
    </div>
  )
}

function Flow() {
  return (
    <div className="flex-1 mx-4 h-[2px] bg-gray-800 relative overflow-hidden">
      <div className="absolute h-full w-1/3 bg-blue-500 animate-flow" />

      <style jsx>{`
        @keyframes flow {
          0% { transform: translateX(-100%) }
          100% { transform: translateX(300%) }
        }
        .animate-flow {
          animation: flow 1.6s linear infinite;
        }
      `}</style>
    </div>
  )
}