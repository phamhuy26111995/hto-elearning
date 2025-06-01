import { FormCourse } from "@/types/course";
import { FormProvider, useFieldArray, useForm } from "react-hook-form";
import { useParams } from "react-router";
import Module from "./Module";
import { Button } from "@/components/ui/button";

export default function CourseDetail() {
  const { courseId } = useParams();
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

  return (
    <FormProvider {...methods}>
      {moduleFields.map((module, index) => {
        return <Module id={index} key={module.id} />;
      })}
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
