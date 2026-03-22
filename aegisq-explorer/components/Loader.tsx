"use client"

export default function Loader({ label = "Initializing AegisQ Network" }) {

  return (
    <div className="relative flex items-center justify-center h-[75vh] overflow-hidden">

      {/* BACKGROUND GRID */}
      <div className="absolute inset-0 opacity-[0.04] bg-[linear-gradient(#fff_1px,transparent_1px),linear-gradient(90deg,#fff_1px,transparent_1px)] bg-[size:40px_40px]" />

      {/* CENTER CONTENT */}
      <div className="relative flex flex-col items-center gap-8">

        {/* CORE SYSTEM NODE */}
        <div className="relative">

          {/* outer pulse */}
          <div className="w-20 h-20 rounded-full bg-blue-500/10 animate-ping" />

          {/* middle glow */}
          <div className="absolute inset-0 m-auto w-16 h-16 rounded-full bg-blue-500/20 blur-xl" />

          {/* core */}
          <div className="absolute inset-0 flex items-center justify-center">
            <div className="w-6 h-6 border-2 border-blue-400 border-t-transparent rounded-full animate-spin" />
          </div>

        </div>

        {/* TEXT */}
        <div className="text-center space-y-2">

          <p className="text-blue-400 text-xs tracking-[0.3em] uppercase">
            AegisQ System
          </p>

          <p className="text-white text-sm font-medium">
            {label}
          </p>

          {/* animated dots */}
          <div className="flex justify-center gap-1 mt-2">
            <Dot delay="0s" />
            <Dot delay="0.2s" />
            <Dot delay="0.4s" />
          </div>

        </div>

        {/* DATA STREAM LINE */}
        <div className="w-80 h-[2px] bg-gray-800 relative overflow-hidden rounded">

          <div className="absolute h-full w-1/3 bg-blue-500 animate-stream" />

        </div>

        {/* SYSTEM STATUS TEXT */}
        <div className="text-[11px] text-gray-500 tracking-wide">

          <span className="text-green-400">●</span> CONNECTING TO VALIDATORS  
          <span className="mx-2">•</span>
          <span className="text-blue-400">SYNCING STATE</span>

        </div>

      </div>

      {/* FLOATING PARTICLES */}
      <div className="absolute inset-0 pointer-events-none">

        <Particle top="20%" left="15%" delay="0s" />
        <Particle top="70%" left="80%" delay="1s" />
        <Particle top="40%" left="60%" delay="0.5s" />
        <Particle top="80%" left="30%" delay="1.5s" />

      </div>

      {/* ANIMATIONS */}
      <style jsx>{`
        @keyframes stream {
          0% { transform: translateX(-100%) }
          100% { transform: translateX(300%) }
        }

        .animate-stream {
          animation: stream 1.4s linear infinite;
        }
      `}</style>

    </div>
  )
}

/* DOTS */
function Dot({ delay }: { delay: string }) {
  return (
    <span
      className="w-1.5 h-1.5 bg-blue-400 rounded-full animate-bounce"
      style={{ animationDelay: delay }}
    />
  )
}

/* PARTICLES */
function Particle({ top, left, delay }: any) {
  return (
    <span
      className="absolute w-1 h-1 bg-blue-400 rounded-full opacity-30 animate-pulse"
      style={{
        top,
        left,
        animationDelay: delay
      }}
    />
  )
}