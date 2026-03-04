"use client"

import {useEffect,useState} from "react"
import {getBlocks} from "../../lib/api"
import Link from "next/link"

export default function BlocksPage(){

  const [blocks,setBlocks] = useState<any[]>([])

  useEffect(()=>{
    getBlocks().then(setBlocks)
  },[])

  return(

    <div>

      <h1 className="text-3xl mb-6">
        Blocks
      </h1>

      <div className="grid grid-cols-2 gap-4">

        {blocks.map((b)=>(
          
          <Link key={b.height} href={`/block/${b.height}`}>

            <div className="bg-gray-900 p-4 rounded hover:bg-gray-800">

              <div className="text-lg">
                Block {b.height}
              </div>

              <div className="text-gray-400 text-sm">
                Transactions: {b.txs}
              </div>

              <div className="text-gray-500 text-xs mt-2">
                {b.hash.slice(0,25)}...
              </div>

            </div>

          </Link>

        ))}

      </div>

    </div>
  )
}