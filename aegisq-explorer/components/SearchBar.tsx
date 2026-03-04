"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"

export default function SearchBar(){

  const router = useRouter()
  const [query,setQuery] = useState("")

  function search(e:any){

    e.preventDefault()

    if(!query) return

    if(!isNaN(Number(query))){
      router.push(`/block/${query}`)
    }else{
      router.push(`/txhash/${query}`)
    }

  }

  return(

    <form onSubmit={search} className="flex gap-2">

      <input
        className="bg-gray-900 border border-gray-700 p-2 rounded w-80"
        placeholder="Search block height or tx hash"
        value={query}
        onChange={(e)=>setQuery(e.target.value)}
      />

      <button className="bg-blue-600 px-4 rounded">
        Search
      </button>

    </form>

  )
}