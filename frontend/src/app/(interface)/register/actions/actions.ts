
'use client'

import { useState } from 'react'
import axios from 'axios'
import { useRouter } from 'next/navigation'
import {z} from "zod"
import { useForm } from 'react-hook-form'
import { registerSchema, registerType } from '@/lib/validations'
import {zodResolver} from "@hookform/resolvers/zod"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Divide } from 'lucide-react'
 
export default function RegisterForm() {
  const router = useRouter()
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  //react hook RegisterForm
  
    const form = useForm<registerType>({
    resolver : zodResolver(registerSchema),
        })
  async function handleSubmit(data:registerType) {
  console.log(data)
     }


  return (
  <div>
  </div>
  )
}
