"use client"
import { useState } from 'react'
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { DownloadModal } from "../download-modal"

export function Hero() {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [videoUrl, setVideoUrl] = useState('')

  const handleDownload = () => {
    // Here you would implement the actual download logic
    if(videoUrl) {
      setIsModalOpen(true)
    }
  }
  return (
    <div className="text-center space-y-6">
      <h1 className="text-4xl md:text-5xl font-bold">
        Download Videos from Any Platform
      </h1>
      <p className="text-slate-400">
        Easily download videos from YouTube, Facebook, Instagram, and more with just one click
      </p>
      <div className="max-w-xl mx-auto flex gap-2">
      <Input
              placeholder="Paste video link here..."
              className="bg-slate-900 border-slate-800"
              value={videoUrl}
              onChange={(e) => setVideoUrl(e.target.value)}
            />
        <Button className="bg-orange-500 hover:bg-orange-600" onClick={handleDownload}>
          Download
        </Button>
      </div>
      <DownloadModal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} videoUrl={videoUrl} />
    </div>
  )
}
