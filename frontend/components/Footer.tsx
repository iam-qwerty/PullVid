import { Twitter, Facebook, Instagram } from 'lucide-react'

export function Footer() {
  return (
    <footer className="border-t border-slate-800 mt-20">
        <div className="container mx-auto px-4 py-8">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div>
              <h3 className="font-semibold mb-4">VideoGrab</h3>
              <p className="text-sm text-slate-400">
                Download your favorite videos from social media platforms with ease.
              </p>
            </div>
            <div>
              <h3 className="font-semibold mb-4">Quick Links</h3>
              <ul className="space-y-2 text-sm text-slate-400">
                <li>Home</li>
                <li>How It Works</li>
                <li>Supported Sites</li>
                <li>FAQ</li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold mb-4">Legal</h3>
              <ul className="space-y-2 text-sm text-slate-400">
                <li>Privacy Policy</li>
                <li>Terms of Service</li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold mb-4">Follow Us</h3>
              <div className="flex gap-4 text-slate-400">
                <Twitter className="w-5 h-5" />
                <Facebook className="w-5 h-5" />
                <Instagram className="w-5 h-5" />
              </div>
            </div>
          </div>
          <div className="mt-8 pt-8 border-t border-slate-800 text-center text-sm text-slate-400">
            © 2023 VideoGrab. All rights reserved.
          </div>
        </div>
      </footer>
  )
}
