import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
  } from "@/components/ui/accordion"
  
  export function FAQ() {
    const faqItems = [
      {
        question: "Is it legal to download videos from social media platforms?",
        answer: "The legality of downloading videos from social media platforms can vary depending on the platform's terms of service and copyright laws in your country. Generally, it's legal to download videos for personal use, but redistributing or using them commercially may infringe on copyright laws.",
      },
      {
        question: "What video formats are supported?",
        answer: "We support a wide range of video formats, including MP4, WebM, AVI, and MOV. The available formats may vary depending on the source platform and the original video format.",
      },
      {
        question: "Is there a limit to how many videos I can download?",
        answer: "There is no strict limit on the number of videos you can download. However, we encourage responsible use of our service and adherence to the terms of service of the source platforms.",
      },
      {
        question: "Can I download private or restricted videos?",
        answer: "No, our service can only download publicly available videos. Private or restricted videos require authentication and cannot be accessed by our downloader.",
      },
      {
        question: "Is this service free to use?",
        answer: "Yes, our basic service is free to use. We may offer premium features or higher download speeds for a fee in the future.",
      },
    ]
  
    return (
      <div className="mt-20 max-w-3xl mx-auto" id="faq">
        <h2 className="text-2xl font-semibold mb-4 text-center">Frequently Asked Questions</h2>
        <Accordion type="single" collapsible className="w-full">
          {faqItems.map((item, index) => (
            <AccordionItem key={index} value={`item-${index}`}>
              <AccordionTrigger>{item.question}</AccordionTrigger>
              <AccordionContent>{item.answer}</AccordionContent>
            </AccordionItem>
          ))}
        </Accordion>
      </div>
    )
  }
  
  