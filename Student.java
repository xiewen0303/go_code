package dbbean;

import java.io.Serializable;
import com.kernel.data.dao.AbsVersion;
import com.kernel.data.dao.IEntity;
import com.kernel.check.db.annotation.*;
import java.sql.*;

@Table("student")
public class  extends AbsVersion implements Serializable,IEntity{

    @EntityField
	private static final long serialVersionUID = 1L;

 
    @Column("age")
    private int age;
 
    @Column("name")
    private String name;
 

}