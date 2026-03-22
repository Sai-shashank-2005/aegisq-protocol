"use client"

import Link from "next/link"
import { usePathname } from "next/navigation"
import Image from "next/image"
import {
  LayoutDashboard,
  Blocks,
  ArrowRightLeft,
  Activity
} from "lucide-react"

export default function Sidebar() {

  const pathname = usePathname()

  return (

    <div className="
      w-72 h-screen fixed left-0 top-0 z-50
      bg-gradient-to-b from-gray-950 to-black
      border-r border-white/5
      flex flex-col justify-between
    ">

      {/* ===== BRAND ===== */}
      <div>

        <div className="px-6 py-6 border-b border-white/5">

          <div className="flex items-center gap-3">

            <div className="relative">

              <Image
                src="/logov2.png"
                width={42}
                height={42}
                alt="AegisQ"
                className="rounded-lg"
              />

              {/* subtle glow */}
              <div className="absolute inset-0 rounded-lg bg-blue-500/10 blur-md opacity-40" />

            </div>

            <div>
              <h1 className="text-white font-semibold tracking-tight">
                AegisQ Explorer
              </h1>
              <p className="text-xs text-gray-500">
                Hybrid BFT Blockchain
              </p>
            </div>

          </div>

        </div>

        {/* ===== NAV ===== */}
        <nav className="px-3 py-6 space-y-2">

          <SidebarLink
            href="/"
            label="Dashboard"
            icon={<LayoutDashboard size={18} />}
            active={pathname === "/"}
          />

          <SidebarLink
            href="/blocks"
            label="Blocks"
            icon={<Blocks size={18} />}
            active={pathname === "/blocks"}
          />

          {/* ❌ REMOVE IF NOT IMPLEMENTED */}
          {/* 
          <SidebarLink
            href="/transactions"
            label="Transactions"
            icon={<ArrowRightLeft size={18} />}
            active={pathname === "/transactions"}
          />

          <SidebarLink
            href="/network"
            label="Network"
            icon={<Activity size={18} />}
            active={pathname === "/network"}
          />
          */}

        </nav>

      </div>

      {/* ===== FOOTER ===== */}
      <div className="px-5 py-5 border-t border-white/5">

        <div className="flex items-center gap-2 text-xs text-gray-400">

          <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></span>

          Network Operational

        </div>

        <div className="mt-3 text-[11px] text-gray-600 space-y-1">

          <p>Validators: 4</p>
          <p>Fault Tolerance: f = 1</p>
          <p>Consensus: BFT</p>

        </div>

        <p className="text-[10px] text-gray-700 mt-4">
          AegisQ Node v1.0
        </p>

      </div>

    </div>

  )
}

function SidebarLink({ href, label, icon, active }: any) {

  return (

    <Link
      href={href}
      className={`
        group relative flex items-center gap-3
        px-4 py-3 rounded-xl text-sm
        transition-all duration-300 overflow-hidden

        ${active
          ? "bg-gradient-to-r from-blue-500/20 to-blue-500/5 text-blue-400 border border-blue-500/30"
          : "text-gray-400 hover:text-white hover:bg-white/5"
        }
      `}
    >

      {/* glow hover */}
      <div className="
        absolute inset-0 opacity-0 group-hover:opacity-100
        bg-gradient-to-r from-blue-500/10 to-transparent
        transition duration-300
      " />

      {/* left indicator */}
      {active && (
        <span className="
          absolute left-0 top-2 bottom-2 w-[2px]
          bg-blue-400 rounded-r shadow-[0_0_8px_rgba(59,130,246,0.8)]
        " />
      )}

      <div className="relative z-10 group-hover:scale-110 transition">
        {icon}
      </div>

      <span className="relative z-10">{label}</span>

    </Link>

  )
}