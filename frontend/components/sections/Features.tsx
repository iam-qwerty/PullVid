import { Zap, Shield, Infinity } from 'lucide-react'

const features = [
    {
      title: "Fast Download",
      description: "Download videos at maximum speed with our optimized servers",
      icon: <Zap className="w-5 h-5 text-orange-500" />,
    },
    {
      title: "Secure & Safe",
      description: "100% safe downloads without any harmful viruses or ads",
      icon: <Shield className="w-5 h-5 text-orange-500" />,
    },
    {
      title: "Unlimited Downloads",
      description: "No limits on the number of videos you can download",
      icon: <Infinity className="w-5 h-5 text-orange-500" />,
    },
  ]

export function Features() {
  return (
    <div className="mt-20 grid md:grid-cols-3 gap-8 max-w-4xl mx-auto">
      {features.map((feature) => (
        <div
          key={feature.title}
          className="bg-slate-900 p-6 rounded-lg space-y-2"
        >
          <div className="w-10 h-10 rounded-lg bg-orange-500/10 flex items-center justify-center">
            {feature.icon}
          </div>
          <h3 className="font-semibold">{feature.title}</h3>
          <p className="text-sm text-slate-400">{feature.description}</p>
        </div>
      ))}
    </div>
  )
}
