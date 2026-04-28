import { Hero } from "@/components/sections/Hero"
import { Platforms } from "@/components/sections/Platforms"
import { Features } from "@/components/sections/Features"
import { HowItWorks } from "@/components/sections/How-It-Works"
import { FAQ } from "@/components/sections/FAQ"

export default function VideoDownloader() {
  return (
    <div className="container mx-auto px-4 py-16">
      <Hero />
      <Platforms />
      <HowItWorks/>
      <Features />
      <FAQ />
    </div>
  )
}





