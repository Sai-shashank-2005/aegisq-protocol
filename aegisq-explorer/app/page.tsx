"use client"

import {useEffect,useState} from "react"
import {getStatus,getBlocks} from "../lib/api"
import Link from "next/link"

export default function Home(){

  const [status,setStatus] = useState<any>(null)
  const [blocks,setBlocks] = useState<any[]>([])

  useEffect(()=>{
    getStatus().then(setStatus)
    getBlocks().then(setBlocks)
  },[])

  return(

    <div>

      <h1 className="text-3xl mb-6">
        AegisQ Network
      </h1>

      {status && (
        <div className="bg-gray-900 p-4 rounded mb-6">
          Latest Height: {status.height}
        </div>
      )}

      <h2 className="text-xl mb-4">
        Recent Blocks
      </h2>

      <div className="space-y-2">

        {blocks.slice(0,10).map((b)=>(
          
          <Link key={b.height} href={`/block/${b.height}`}>

            <div className="bg-gray-900 p-3 rounded hover:bg-gray-800">

              Block {b.height} • {b.txs} tx

            </div>

          </Link>

        ))}

      </div>

    </div>
  )
}