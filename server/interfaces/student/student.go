package student

import (
	"context"
	"database/sql"

	f "fmt"
	i "grpctest/server/interfaces"

	_ "github.com/go-sql-driver/mysql"

	model "grpctest/server/models"
	pb "grpctest/server/proto"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tesgrpc")
	if err != nil {
		return nil
	}
	// defer db.Close()

	return db

}

// for abstraction package

func NewStudentInterfaces() i.StudentRepo {
	return &Student{
		DB: connect(),
	}
}

type Student struct {
	DB *sql.DB
}

func (v *Student) Create(ctx context.Context, in *pb.Student) (*pb.StudentID, error) {
	sql := "insert into dbgrpc values(?,?,?)"
	stmt, err := v.DB.Prepare(sql)
	if err != nil {
		f.Println("stmt err", err)
	}

	defer stmt.Close()

	data := model.Student{}
	data.Id = in.Id
	data.Name = in.Name
	data.Grade = in.Grade

	stmt.Exec(data.Id, data.Name, data.Grade)

	if err != nil {
		f.Println(err.Error())
		return nil, err
	}

	f.Println("berhasil masuk SQL", in.Id, in.Name, in.Grade)
	v.DB.Prepare("")

	return nil, nil
}

func (v *Student) Read(ctx context.Context, in *pb.StudentID) (*pb.Student, error) {
	sql := "select id, name, grade from dbgrpc"
	rows, err := v.DB.Query(sql)

	if err != nil {
		f.Println(err.Error())
	}
	defer rows.Close()

	var students = model.Student{}

	for rows.Next() {
		var Id int64
		var Name string
		var Grade int32

		// var stu = &model.Student
		err := rows.Scan(&Id, &Name, &Grade)
		// students = append(students, stu)

		if err != nil {
			f.Println(err.Error())
		}
		students.Id = Id
		students.Name = Name
		students.Grade = Grade
	}

	payload := &pb.Student{}
	payload.Id = students.Id
	payload.Name = students.Name
	payload.Grade = students.Grade

	return payload, nil
}

func (v *Student) Update(ctx context.Context, in *pb.Student) (*pb.StudentID, error) {
	sql := "update dbgrpc set name = ?, grade = ? where id = ?"
	stmt, err := v.DB.Prepare(sql)
	if err != nil {
		f.Println("stmt err", err)
	}
	defer stmt.Close()

	data := model.Student{}
	data.Id = in.Id
	data.Name = in.Name
	data.Grade = in.Grade

	stmt.Exec(data.Name, data.Grade, data.Id)

	if err != nil {
		f.Println(err.Error())
		return nil, err
	}

	return nil, nil
}

func (v *Student) Delete(ctx context.Context, in *pb.StudentID) (*pb.StudentID, error) {
	sql := "delete from dbgrpc where id=?"
	stmt, err := v.DB.Prepare(sql)
	if err != nil {
		f.Println(err)
	}
	defer stmt.Close()

	stmt.Exec(in.Id)

	if err != nil {
		return nil, err
	}

	// fmt.Println(&in.Name)
	return nil, nil
}
