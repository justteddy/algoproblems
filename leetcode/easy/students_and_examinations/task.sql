select t1.student_id, t1.student_name, t1.subject_name, count(t2.student_id) as attended_exams
from(
        select s.student_id, s.student_name, su.subject_name
        from Students s, Subjects su
    ) t1 left join Examinations t2 on t1.student_id = t2.student_id and t1.subject_name=t2.subject_name
group by t1.student_id, t1.subject_name
order by student_id;
