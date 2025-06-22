import { FormCourse } from "@/types/course";
import { FormProvider, set, useFieldArray, useForm } from "react-hook-form";
import { useParams } from "react-router";
import Module from "./Module";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { useState } from "react";
import { Input } from "@/components/ui/input";

export default function CourseDetail() {
  const { courseId } = useParams();
  const [tabIndex, setTabIndex] = useState<number>(0);
  const methods = useForm<FormCourse>({
    defaultValues: {
      modules: [
        {
          title: "",
          description: "",
          lessons: [],
          quizzes: [],
          orderIndex: 0,
        },
      ],
    },
  });

  const {
    control,
    formState: { errors },
    handleSubmit,
  } = methods;

  const {
    fields: moduleFields,
    append: appendModule,
    remove: removeModule,
  } = useFieldArray<FormCourse, "modules">({
    control,
    name: "modules",
  });

  function onSubmit(data: FormCourse) {
    console.log(data);
  }

  function onChangeGoToTab(e: any) {
    const value = e.target.value;
    if (value < 0 || isNaN(value) || value === "") {
      setTabIndex(0);
      return;
    }

    const tab = +value - 1;

    if (tab >= moduleFields.length) {
      setTabIndex(moduleFields.length - 1);
      return;
    }

    setTabIndex(tab);
  }

  const tabNum = +tabIndex + 1;

  return (
    <FormProvider {...methods}>
      <div className="flex gap-3 flex-col p-4">
        <span className="flex gap-2 items-center">
          <span className="text-[1.1rem]">Go to tab : </span>
          <Input
            min={1}
            maxLength={moduleFields.length}
            value={tabNum}
            onChange={onChangeGoToTab}
            style={{ maxWidth: "100px" }}
            type="number"
          />
        </span>
        <Tabs
          defaultValue={"0"}
          value={tabIndex.toString()}
          onValueChange={(value) => setTabIndex(Number(value))}
          className="w-[400px]"
        >
          <TabsList>
            {moduleFields.map((module, index) => (
              <TabsTrigger key={module.id} value={index.toString()}>
                Module {index + 1}
              </TabsTrigger>
            ))}
          </TabsList>
          {moduleFields.map((module, index) => (
            <TabsContent key={module.id} value={index.toString()}>
              <Module id={index} />
            </TabsContent>
          ))}
        </Tabs>
      </div>
      <Button
        onClick={() =>
          appendModule({
            title: "",
            description: "",
            lessons: [],
            quizzes: [],
            orderIndex: 0,
            courseId: Number(courseId),
            moduleId: 0,
          })
        }
      >
        Add Module
      </Button>

      <Button onClick={handleSubmit(onSubmit)}>Submit Form</Button>
    </FormProvider>
  );
}
