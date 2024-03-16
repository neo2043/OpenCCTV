"use client"

import { useRef, useState } from "react"

export default function PlusSign(){
    const dialogRef = useRef<HTMLDialogElement>(null)

    return (
        <div className="rounded-full">
            <DialogElement dialogRef={dialogRef} />
            <button onClick={() => {
                dialogRef.current?.showModal()
            }}>Click to open</button>
        </div>
    )
}

const DialogElement = ({dialogRef} : {
    dialogRef: any
}) => {
    const [state, setState] = useState("")
    return <dialog onClose={() => {
        setState("")
    }} ref={dialogRef}>
         <button onClick={() => {
                dialogRef.current?.close()
            }}>Click to close</button>
            <h1>{}</h1>
        <div></div>
    </dialog>
}