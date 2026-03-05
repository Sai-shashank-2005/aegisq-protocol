"use client"

import SearchBar from "./SearchBar"

export default function Header(){

return(

<div className="h-16 border-b border-gray-800 flex items-center justify-between px-8 bg-black">

<h2 className="text-sm text-gray-400">
AegisQ Blockchain Explorer
</h2>

<SearchBar/>

</div>

)

}