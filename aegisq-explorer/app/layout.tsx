import "./globals.css"
import Sidebar from "@/components/SideBar"
import Header from "@/components/Header"

export default function RootLayout({children}:{children:React.ReactNode}){

return(

<html lang="en">

<body className="bg-black text-white">

<Sidebar/>

<div className="ml-72">

<Header/>

<div className="p-8">
{children}
</div>

</div>

</body>

</html>

)

}