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

export default function Sidebar(){

const pathname = usePathname()

return(

<div className="w-72 h-screen bg-gradient-to-b from-gray-950 to-gray-900 border-r border-gray-800 flex flex-col justify-between fixed left-0 top-0">

{/* ===== BRAND ===== */}

<div>

<div className="px-6 py-6 border-b border-gray-800">

<div className="flex items-center gap-3">

<Image
src="/logo.png"
width={42}
height={42}
alt="AegisQ"
/>

<div>

<h1 className="text-lg font-semibold text-white">
AegisQ Explorer
</h1>

<p className="text-xs text-gray-500">
Hybrid BFT Blockchain
</p>

</div>

</div>

</div>

{/* ===== NAVIGATION ===== */}

<nav className="px-4 py-6 space-y-2">

<SidebarLink
href="/"
label="Dashboard"
icon={<LayoutDashboard size={18}/>}
active={pathname === "/"}
/>

<SidebarLink
href="/blocks"
label="Blocks"
icon={<Blocks size={18}/>}
active={pathname === "/blocks"}
/>

<SidebarLink
href="/transactions"
label="Transactions"
icon={<ArrowRightLeft size={18}/>}
active={pathname === "/transactions"}
/>

<SidebarLink
href="/network"
label="Network"
icon={<Activity size={18}/>}
active={pathname === "/network"}
/>

</nav>

</div>

{/* ===== FOOTER ===== */}

<div className="px-6 py-5 border-t border-gray-800">

<div className="flex items-center gap-2 text-xs text-gray-500">

<span className="w-2 h-2 bg-green-500 rounded-full animate-pulse"></span>

Network Operational

</div>

<p className="text-xs text-gray-600 mt-2">
AegisQ Node v1.0
</p>

</div>

</div>

)

}

function SidebarLink({href,label,icon,active}:any){

return(

<Link
href={href}
className={`group relative flex items-center gap-3 px-4 py-3 rounded-xl text-sm font-medium transition-all duration-200
${active
? "bg-blue-500/10 text-blue-400"
: "text-gray-400 hover:bg-gray-800 hover:text-white"
}`}

>

{active && (
<span className="absolute left-0 top-2 bottom-2 w-1 bg-blue-500 rounded-r"></span>
)}

{icon}

{label}

</Link>

)

}