import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";

export default function TestingPage() {
 const {setValue , control} = useForm<any>({
    defaultValues : "default"
  })




  return (
    <div>
      <Controller 
        control={control}
        name="radio"
        render={({ field }) => <RadioGroupDemo {...field} />}
      />
      <Button onClick={() => setValue("radio", "comfortable")}>Set Value for Radio</Button>
    </div>
  );
}

function RadioGroupDemo({ value, onChange }: any) {
  return (
    <RadioGroup
      value={value}
      onValueChange={onChange}

    >
      <div className="flex items-center gap-3">
        <RadioGroupItem value="default" id="r1" />
        <Label htmlFor="r1">Default</Label>
      </div>
      <div className="flex items-center gap-3">
        <RadioGroupItem value="comfortable" id="r2" />
        <Label htmlFor="r2">Comfortable</Label>
      </div>
      <div className="flex items-center gap-3">
        <RadioGroupItem value="compact" id="r3" />
        <Label htmlFor="r3">Compact</Label>
      </div>
    </RadioGroup>
  );
}
