import "./globals.css"
import Navbar from "@/components/Navbar"

export default function RootLayout({children}:{children:React.ReactNode}){

return(

<html>
<body className="bg-black text-white">

<Navbar/>

<div className="max-w-6xl mx-auto py-10">
{children}
</div>

</body>
</html>

)
}