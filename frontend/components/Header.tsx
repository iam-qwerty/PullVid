'use client'

import Link from 'next/link'
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"
import { Button } from "@/components/ui/button"
import { Menu } from 'lucide-react'
import Image from 'next/image'

export function Header() {
  const navigation = [
    ['Home', '/'],
    ['How It Works', '#how-it-works'],
    ['Supported Sites', '#platforms'],
    ['FAQ', '#faq'],
  ]

  return (
    <header className="bg-slate-950 border-b border-slate-800">
      <div className="container mx-auto px-4">
        <div className="flex h-16 items-center justify-between">
          <Link href="/" className="flex items-center gap-2">
          <Image
            src={'/logo.svg'}
            alt='logo'
            width={40}
            height={50}
          />
            {/* <svg
              width="32"
              height="32"
              viewBox="0 0 32 32"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              className="text-orange-500"
            >
              <circle cx="16" cy="16" r="16" fill="currentColor"/>
              <path
                d="M22 16L13 21.1962L13 10.8038L22 16Z"
                fill="white"
              />
            </svg> */}
            <span className="text-white text-xl font-semibold">PullVid</span>
          </Link>
          
          <nav className="hidden md:flex gap-8">
            {navigation.map(([label, href]) => (
              <Link
                key={label}
                href={href}
                className="text-slate-200 hover:text-white transition-colors"
              >
                {label}
              </Link>
            ))}
          </nav>

          <Sheet>
            <SheetTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="md:hidden text-slate-200 hover:text-white"
              >
                <Menu className="h-6 w-6" />
                <span className="sr-only">Toggle menu</span>
              </Button>
            </SheetTrigger>
            <SheetContent side="right" className="w-[300px] bg-slate-950 border-slate-800 text-white">
              <nav className="flex flex-col gap-4 mt-8">
                {navigation.map(([label, href]) => (
                  <Link
                    key={label}
                    href={href}
                    className="text-slate-200 hover:text-white transition-colors text-lg"
                  >
                    {label}
                  </Link>
                ))}
              </nav>
            </SheetContent>
          </Sheet>
        </div>
      </div>
    </header>
  )
}

