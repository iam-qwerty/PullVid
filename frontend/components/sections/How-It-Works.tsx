import { Link, Download, Check } from 'lucide-react'

export function HowItWorks() {
  const steps = [
    {
      icon: <Link className="w-8 h-8 text-orange-500" />,
      title: "Copy the Link",
      description: "Copy the video URL from your favorite platform",
    },
    {
      icon: <Download className="w-8 h-8 text-orange-500" />,
      title: "Paste and Download",
      description: "Paste the link, select the format and quality, and start downloading",
    },
    {
      icon: <Check className="w-8 h-8 text-orange-500" />,
      title: "Enjoy Your Video",
      description: "Your video is ready to watch offline",
    },
  ]

  return (
    <div className="mt-20 text-center" id="how-it-works">
      <h2 className="text-2xl font-semibold mb-4">How It Works</h2>
      <div className="grid md:grid-cols-3 gap-8 max-w-4xl mx-auto mt-8">
        {steps.map((step, index) => (
          <div key={index} className="flex flex-col items-center">
            <div className="w-16 h-16 rounded-full bg-slate-800 flex items-center justify-center mb-4">
              {step.icon}
            </div>
            <h3 className="text-lg font-semibold mb-2">{step.title}</h3>
            <p className="text-slate-400">{step.description}</p>
          </div>
        ))}
      </div>
    </div>
  )
}

