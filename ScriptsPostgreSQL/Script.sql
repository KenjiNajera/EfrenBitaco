-- 
-- TABLE: Members 
--trigeer insert
-- userID no se puede repetir
-- al reasignar a otro usuario hay que cambiar el id del proyecto un update
-- al concluir un proyecto hay que eliminar a los miembros de esta tabla ---------------------------------------------------------------------------------------------------
CREATE TABLE Members(
    MemberID   Serial NOT NULL,
    UserID     int4    NOT NULL,
    ProjectID  int4    NOT NULL,
    CONSTRAINT "PK_MemberID" PRIMARY KEY (MemberID),
	CONSTRAINT "UQ_UserIDProjectID" UNIQUE(UserID,ProjectID)
);
	
-- 
-- TABLE: Binnacle 
--
CREATE TABLE Binnacles(
    BinnacleID   Serial         NOT NULL,
	MonthName	 varchar(15)    NOT NULL,
    Month        int4           NOT NULL,
    Year         int4           NOT NULL,
    TotalHours   decimal        ,
	CreateAt	 date           NOT NULL default current_date,
	Status       boolean        NOT NULL default false,
    UserID       int4           NOT NULL,
    CONSTRAINT "PK_BinnacleID" PRIMARY KEY (BinnacleID),
	CONSTRAINT "UQ_yearMonth" UNIQUE(Year,Month,UserID),
	CONSTRAINT "CK_ValidYear" CHECK(Year = extract(Year from current_date))
);

-- 
-- TABLE: Activitys 
--
CREATE TABLE Activitys(
    ActivityID  Serial  		  NOT NULL,
    Summary     varchar(200)      NOT NULL,
    Invoice     varchar(10)       NOT NULL default '',
    DayID       int4 			  NOT NULL,
    CONSTRAINT "PK_ActivityID" PRIMARY KEY (ActivityID),
	CONSTRAINT "UQ_DayIDActivities" UNIQUE (DayID)
);

-- 
-- TABLE: Certificates
--
CREATE TABLE Certificates(
    CertificateID  Serial  NOT NULL,
    IDCertificate   varchar(30)    NOT NULL,
    CertificateName  varchar(25)   NOT NULL,
    Expedition      varchar(15)    NOT NULL,
	Expiration      varchar(20)    NOT NULL default '(without caducity)',
    ResourceDataID  int4           NOT NULL,
    CONSTRAINT "PK_CertificateID" PRIMARY KEY (CertificateID),
	CONSTRAINT "UQ_Certificate" UNIQUE(IDCertificate)
);

-- 
-- TABLE: CheckTime 
--
CREATE TABLE CheckTimes(
    CheckTimeID    Serial NOT NULL,
    CheckTime      time         NOT NULL default '00:00',
    CreateAt	   timestamp    NOT NULL default current_timestamp,
    DayID          int4         NOT NULL,
    CONSTRAINT "PK_CheckTimeID" PRIMARY KEY (CheckTimeID),
	CONSTRAINT "UQ_DayIDCheck" UNIQUE(DayID)
);

-- 
-- TABLE: Client 
--
CREATE TABLE Clients(
    ClientID      Serial  NOT NULL,
    ClientName    varchar(45)      NOT NULL,
    SocialReason  varchar(200)     NOT NULL,
    PhotoClient   varchar(1000)    NOT NULL default '',
    RFC           varchar(20)      NOT NULL,-- el rfc es difente al tener el mismo cliente con diferente razon social?
    Address       varchar(100),
    City          varchar(20),
    CP            varchar(10),
    CellPhone     varchar(12),
    Status        boolean          NOT NULL default true,
    CountryID     int4 NOT NULL default 0,
    CONSTRAINT "PK_ClientID" PRIMARY KEY (ClientID)
);

-- 
-- TABLE: Country 
--
CREATE TABLE Countrys(
    CountryID      Serial 			NOT NULL,
   	CountryName    varchar(45)      	NOT NULL,
    CONSTRAINT "PK_CountryID" PRIMARY KEY (CountryID),
	CONSTRAINT "UQ_CountryName" UNIQUE(CountryName)
);

-- 
-- TABLE: Day 
--
CREATE TABLE Days(
    DayID       Serial NOT NULL,
    DayDate     date           NOT NULL,
	CreateAt    timestamp      NOT NULL default current_timestamp,
	TotalHoursDay time,
    DailyTime   time,
    BinnacleID  int4           NOT NULL,
    CONSTRAINT "PK_DayID" PRIMARY KEY (DayID),
	CONSTRAINT "UQ_DayDate" UNIQUE(DayDate, BinnacleID),
	CONSTRAINT "CK_DayDate_CreateAt" CHECK(DayDate < CreateAt)
);

CREATE TABLE DayProjects(
	DayprojectID Serial    NOT NULL,
	DayID        int4      NOT NULL,
    ProjectID    int4  	   NOT NULL,
    CONSTRAINT "PK_DayprojectID" PRIMARY KEY (DayprojectID),
	CONSTRAINT "UQ_DayIDProjectID" UNIQUE(DayID, ProjectID)
	
);

-- 
-- TABLE: DepartureTime 
--
CREATE TABLE DepartureTimes(
    DepartureTimeID    	Serial 		 NOT NULL,
    DepartureTime      	time         NOT NULL default '00:00',
    CreateAt  			timestamp    NOT NULL  default current_timestamp,
    DayID              	int4         NOT NULL,
    CONSTRAINT "PK_DepartureTimeID" PRIMARY KEY (DepartureTimeID),
	CONSTRAINT "UQ_DayIDDeparture" UNIQUE(DayID)
);

-- 
-- TABLE: Education 
--
CREATE TABLE Educations(
    EducationID            Serial         NOT NULL,
	AcademicName           varchar(100)   NOT NULL,
    AcademicTitle          varchar(50)    NOT NULL,
    Year                   int4           NOT NULL,
    ResourceDataID         int4           NOT NULL,
    CONSTRAINT "PK_EducationID" PRIMARY KEY (EducationID),
	CONSTRAINT "CK_YearEducation" CHECK (Year <= extract(Year from current_date)),
	CONSTRAINT "UQ_AcademicResource" UNIQUE(AcademicName, ResourceDataID, AcademicTitle)
);

-- 
-- TABLE: Experience 
--
CREATE TABLE Experiences(
    ExperienceID     Serial          NOT NULL,
    ProjectTitle     varchar(30)     NOT NULL,
    Description      varchar(155)    NOT NULL,
    YearElaboration  int4            NOT NULL,
    ResourceDataID   int4            NOT NULL,
    CONSTRAINT "PK_ExperienceID" PRIMARY KEY (ExperienceID),
	CONSTRAINT "CK_YearExperience" CHECK (YearElaboration <= extract(Year from current_date)),
	CONSTRAINT "UQ_Experiences" UNIQUE(ProjectTitle, ResourceDataID)
);

-- 
-- TABLE: HardSkill 
--
CREATE TABLE HardSkills(
    HardSkillID   Serial   NOT NULL,
    HardSkillName  varchar(25)      NOT NULL,
    Image          varchar(1000)    NOT NULL default '',
    CONSTRAINT "PK_HardSkillID" PRIMARY KEY (HardSkillID),
	CONSTRAINT "UQ_HardSkillName" UNIQUE(HardSkillName)
);

-- 
-- TABLE: HardsResource 
--
CREATE TABLE HardsResources(
    HardsResourceID  Serial  NOT NULL,
    Domain           int4    NOT NULL default 0,
    ResourceDataID   int4    NOT NULL,
    HardSkillID     int4    NOT NULL,
    CONSTRAINT "PK_HardsResourceID" PRIMARY KEY (HardsResourceID),
	CONSTRAINT "CK_DomainHards" CHECK (Domain > 0 and Domain <= 100),
	CONSTRAINT "UQ_HardSkillResource" UNIQUE(ResourceDataID, HardSkillID)
);

-- 
-- TABLE: Languages 
--
CREATE TABLE Languages(
    LanguageID    Serial NOT NULL,
    LanguageName  varchar(20)    NOT NULL,
    CONSTRAINT "PK_LanguageID" PRIMARY KEY (LanguageID),
	CONSTRAINT "UQ_LanguageName" UNIQUE(LanguageName)
);

-- 
-- TABLE: LanguageUser 
--
CREATE TABLE LanguageResources(
    LanguageResourceID  Serial 	NOT NULL,
    Domain          	int4    NOT NULL default 0,
    LanguageID      	int4    NOT NULL,
    ResourceDataID  	int4    NOT NULL,
    CONSTRAINT "PK_LanguageUserID" PRIMARY KEY (LanguageResourceID),
	CONSTRAINT "CK_DomainLanguage" CHECK (Domain >= 0 and Domain <= 100),
	CONSTRAINT "UQ_LanguageResource" UNIQUE(ResourceDataID, LanguageID)
);

-- 
-- TABLE: MealTime 
--
CREATE TABLE MealTimes(
    MealTimeID    	Serial 		 NOT NULL,
    MealTime      	time         NOT NULL default '00:00',
    CreateAt  		timestamp    NOT NULL default current_timestamp,
    DayID         	int4         NOT NULL,
    CONSTRAINT "PK_MealTimeID" PRIMARY KEY (MealTimeID),
	CONSTRAINT "UQ_DayIDMeal" UNIQUE(DayID)
);

-- 
-- TABLE: ModPermiModRol 
--
CREATE TABLE ModPermiModRols(
    ModPermiModRolID   	Serial   		NOT NULL,
    ModulePermissionID  int4            NOT NULL,
    ModuleRolID        	int4            NOT NULL,
    CONSTRAINT "PK_ModPermiModRolID" PRIMARY KEY (ModPermiModRolID)
);

-- 
-- TABLE: ModulePerimission 
--
CREATE TABLE ModulePermissions(
    ModulePermissionID    Serial  		  NOT NULL,
    ModuleID              int4            NOT NULL,
    PermissionID          int4            NOT NULL,
    CONSTRAINT "PK_ModulePermissionID" PRIMARY KEY (ModulePermissionID),
	CONSTRAINT "UQ_ModuleIDPermissionID" UNIQUE(ModuleID,PermissionID)
);

-- 
-- TABLE: ModuleRol 
--
CREATE TABLE ModuleRols(
    ModuleRolID   	Serial  	   NOT NULL,
    RoleID         	int4           NOT NULL,
    ModuleID       	int4           NOT NULL,
    CONSTRAINT "PK_ModuleRolID" PRIMARY KEY (ModuleRolID),
	CONSTRAINT "UQ_ModuleIDRoleID" UNIQUE(ModuleID,RoleID)
);

-- 
-- TABLE: Modules 
--
CREATE TABLE Modules(
    ModuleID    Serial 	 	   NOT NULL,
    ModuleName  varchar(30)    NOT NULL,
    CONSTRAINT "PK_ModuleID" PRIMARY KEY (ModuleID),
	CONSTRAINT "UQ_ModuleName" UNIQUE(ModuleName)
);

-- 
-- TABLE: Overtime 
--
CREATE TABLE Overtimes(
    OvertimeID    Serial 		NOT NULL,
    Overtime      time         	NOT NULL default '00:00',
    CreateAt  	  timestamp    	NOT NULL default current_timestamp,
    DayID         int4         	NOT NULL,
    CONSTRAINT "PK_OvertimeID" PRIMARY KEY (OvertimeID),
	CONSTRAINT "UQ_DayIDovertime" UNIQUE(DayID)
);

-- 
-- TABLE: Permissions 
--
CREATE TABLE Permissions(
    PermissionID    Serial  	  	NOT NULL,
    PermissionName  varchar(50)    	NOT NULL,
    CONSTRAINT "PK_PermissionID" PRIMARY KEY (PermissionID),
	CONSTRAINT "UQ_PermissionName" UNIQUE(PermissionName)
);

-- 
-- TABLE: Projects 
--
CREATE TABLE Projects(
    ProjectID    Serial  		 NOT NULL,
    ProjectName  varchar(30)     NOT NULL,
    Description  varchar(155)    NOT NULL default '',
	CreateAt	 date            NOT NULL default current_date,
	FinishAt     date,
    Status       boolean         NOT NULL default false, -- false es en proceso y true finalizado o concluido
    UserID		 int4,
	ClientID     int4            NOT NULL,
    CONSTRAINT "PK_ProjectID" PRIMARY KEY (ProjectID)
);

-- 
-- TABLE: Requests 
--
CREATE TABLE Requests(
    RequestID      Serial   	   NOT NULL,
    Description    varchar(155)    NOT NULL,
    Status         boolean,
	CreateAt	   date            NOT NULL default current_date,
    Details        varchar(155),
    UserID         int4            NOT NULL,
    TypeRequestID  int4            NOT NULL,
    CONSTRAINT "PK_RequestID" PRIMARY KEY (RequestID)
);

-- 
-- TABLE: ResourceData 
--
CREATE TABLE ResourceDatas(
    ResourceDataID  Serial  		   NOT NULL,
    FirstName       varchar(40)      NOT NULL,
    PhotoUser       varchar(1000)    NOT NULL default '',
    LastName        varchar(40)      NOT NULL,
    About           varchar(155),
    CellPhone       varchar(12)      NOT NULL,
    PersonalEmail   varchar(60)      NOT NULL,
    DateBirth       date             NOT NULL,
    Address         varchar(60)      NOT NULL default '',
    Country         varchar(20)      NOT NULL default '',
    City            varchar(20)	     NOT NULL default '',
    CP              varchar(10),
    CONSTRAINT "PK_ResourceDataID" PRIMARY KEY (ResourceDataID),
	CONSTRAINT "UQ_PersonalEmail" UNIQUE(PersonalEmail)
);

-- 
-- TABLE: Roles 
--
CREATE TABLE Roles(
    RoleID    Serial  NOT NULL,
    RoleName  varchar(25)    NOT NULL,
	Leader    boolean NOT NULL default true,
    CONSTRAINT "PK_RoleID" PRIMARY KEY (RoleID),
	CONSTRAINT "UQ_RoleName" UNIQUE(RoleName)
);

-- 
-- TABLE: SoftSkill 
--
CREATE TABLE SoftSkills(
    SoftSkillID   Serial 		  NOT NULL,
    SoftSkillName  varchar(25)    NOT NULL,
    CONSTRAINT "PK_SoftSkillID" PRIMARY KEY (SoftSkillID),
	CONSTRAINT "UQ_SoftSkillName" UNIQUE(SoftSkillName)
);

-- 
-- TABLE: SoftsResource 
--
CREATE TABLE SoftsResources(
    SoftsResourceID  Serial  NOT NULL,
    SoftSkillID     int4    NOT NULL,
    ResourceDataID   int4    NOT NULL,
    CONSTRAINT "PK_SoftsResourceID" PRIMARY KEY (SoftsResourceID),
	CONSTRAINT "UQ_SoftSkillIDResourceDataID" UNIQUE(SoftSkillID,ResourceDataID)
);

-- 
-- TABLE: TechnologiesProject 
--
CREATE TABLE TechnologiesProjects(
    TechnologiesProjectID  Serial  NOT NULL,
    ProjectID              int4    NOT NULL,
    HardSkillID           int4    NOT NULL,
    CONSTRAINT "PK_TechnologiesProjectID" PRIMARY KEY (TechnologiesProjectID),
	CONSTRAINT "UQ_ProjectIDHardSkillID" UNIQUE(ProjectID,HardSkillID)
);

-- 
-- TABLE: TypeRequest 
--
CREATE TABLE TypeRequests(
    TypeRequestID    Serial  		NOT NULL,
    TypeRequestName  varchar(50)    NOT NULL,
    CONSTRAINT "PK_TypeRequestID" PRIMARY KEY (TypeRequestID),
	CONSTRAINT "UQ_TypeRequestName" UNIQUE(TypeRequestName)
);

-- 
-- TABLE: User 
--
CREATE TABLE Users(
    UserID          Serial           NOT NULL,
    Email           varchar(30)      NOT NULL,
    Password        varchar(1000)   NOT NULL  default '',
    Status          boolean          NOT NULL default false,
	Staterecovery   boolean          NOT NULL default true,
	Mobileid        varchar(200)     NOT NULL default '',
    RoleID          int4             NOT NULL,
    ResourceDataID  int4             NOT NULL,
    CONSTRAINT "PK_UserID" PRIMARY KEY (UserID),
	CONSTRAINT "UQ_ResourceDataIDUSer" UNIQUE(ResourceDataID),
	CONSTRAINT "UQ_Email" UNIQUE(Email)
);

-- 
-- TABLE: Activitys 
--
ALTER TABLE Activitys ADD CONSTRAINT "RefDay73" 
    FOREIGN KEY (DayID)
    REFERENCES Days(DayID) on delete cascade
;

-- 
-- TABLE: Binnacle 
--
ALTER TABLE Binnacles ADD CONSTRAINT "RefUser58" 
    FOREIGN KEY (UserID)
    REFERENCES Users(UserID) on delete cascade
;


-- 
-- TABLE: Certificates
--
ALTER TABLE Certificates ADD CONSTRAINT "RefResourceData45" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade
;

-- 
-- TABLE: CheckTime 
--
ALTER TABLE CheckTimes ADD CONSTRAINT "RefDay71" 
    FOREIGN KEY (DayID)
    REFERENCES Days(DayID) on delete cascade
;

-- 
-- TABLE: Clients 
--
ALTER TABLE Clients ADD CONSTRAINT "RefCountry101" 
    FOREIGN KEY (CountryID)
    REFERENCES Countrys(CountryID) on delete cascade
;

-- 
-- TABLE: Day 
--
ALTER TABLE Days ADD CONSTRAINT "RefBinnacle24" 
    FOREIGN KEY (BinnacleID)
    REFERENCES Binnacles(BinnacleID) on delete cascade
;

-- 
-- TABLE: Dayproject 
--

ALTER TABLE DayProjects ADD CONSTRAINT "RefDay88" 
    FOREIGN KEY (DayID)
    REFERENCES Days(DayID) on delete cascade
;
ALTER TABLE DayProjects ADD CONSTRAINT "RefProject89" 
    FOREIGN KEY (ProjectID)
    REFERENCES Projects(ProjectID) on delete cascade
;

-- 
-- TABLE: DepartureTime 
--
ALTER TABLE DepartureTimes ADD CONSTRAINT "RefDay69" 
    FOREIGN KEY (DayID)
    REFERENCES Days(DayID) on delete cascade
;

-- 
-- TABLE: Education 
--
ALTER TABLE Educations ADD CONSTRAINT "RefResourceData47" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade
;

-- 
-- TABLE: Experience 
--
ALTER TABLE Experiences ADD CONSTRAINT "RefResourceData52" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade
;

-- 
-- TABLE: HardsResource 
--
ALTER TABLE HardsResources ADD CONSTRAINT "RefResourceData13" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade
;

ALTER TABLE HardsResources ADD CONSTRAINT "RefHardSkill86" 
    FOREIGN KEY (HardSkillID)
    REFERENCES HardSkills(HardSkillID) on delete cascade
;

-- 
-- TABLE: LanguageUser 
--
ALTER TABLE LanguageResources ADD CONSTRAINT "RefResourceData37" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade
;

ALTER TABLE LanguageResources ADD CONSTRAINT "RefLanguages38" 
    FOREIGN KEY (LanguageID)
    REFERENCES Languages(LanguageID) on delete cascade
;

-- 
-- TABLE: MealTime 
--

ALTER TABLE MealTimes ADD CONSTRAINT "RefDay70" 
    FOREIGN KEY (DayID)
    REFERENCES Days(DayID) on delete cascade
;

-- 
-- TABLE: Members 
--
ALTER TABLE Members ADD CONSTRAINT "RefProjects76" 
    FOREIGN KEY (ProjectID)
    REFERENCES Projects(ProjectID) on delete cascade
;

ALTER TABLE Members ADD CONSTRAINT "RefUser31" 
    FOREIGN KEY (UserID)
    REFERENCES Users(UserID) on delete cascade
;

-- 
-- TABLE: ModPermiModRol 
--
ALTER TABLE ModPermiModRols ADD CONSTRAINT "RefModulePermission20" 
    FOREIGN KEY (ModulePermissionID)
    REFERENCES ModulePermissions(ModulePermissionID) on delete cascade
;

ALTER TABLE ModPermiModRols ADD CONSTRAINT "RefModuleRol21" 
    FOREIGN KEY (ModuleRolID)
    REFERENCES ModuleRols(ModuleRolID) on delete cascade
;

-- 
-- TABLE: ModulePeimissions 
--
ALTER TABLE ModulePermissions ADD CONSTRAINT "RefPermissions16" 
    FOREIGN KEY (PermissionID)
    REFERENCES Permissions(PermissionID) on delete cascade
;

ALTER TABLE ModulePermissions ADD CONSTRAINT "RefModules17" 
    FOREIGN KEY (ModuleID)
    REFERENCES Modules(ModuleID) on delete cascade
;

-- 
-- TABLE: ModuleRol 
--
ALTER TABLE ModuleRols ADD CONSTRAINT "RefModules18" 
    FOREIGN KEY (ModuleID)
    REFERENCES Modules(ModuleID) on delete cascade
;

ALTER TABLE ModuleRols ADD CONSTRAINT "RefRoles19" 
    FOREIGN KEY (RoleID)
    REFERENCES Roles(RoleID) on delete cascade
;

-- 
-- TABLE: Overtime 
--
ALTER TABLE Overtimes ADD CONSTRAINT "RefDay72" 
    FOREIGN KEY (DayID)
    REFERENCES Days(DayID) on delete cascade
;

-- 
-- TABLE: Projects 
--
ALTER TABLE Projects ADD CONSTRAINT "RefClient9" 
    FOREIGN KEY (ClientID)
    REFERENCES Clients(ClientID) on delete cascade
;
ALTER TABLE Projects ADD CONSTRAINT "RefUser95" 
    FOREIGN KEY (UserID)
    REFERENCES Users(UserID);

-- 
-- TABLE: Requests 
--
ALTER TABLE Requests ADD CONSTRAINT "RefTypeRequest29" 
    FOREIGN KEY (TypeRequestID)
    REFERENCES TypeRequests(TypeRequestID) on delete cascade
;

ALTER TABLE Requests ADD CONSTRAINT "RefUser30" 
    FOREIGN KEY (UserID)
    REFERENCES Users(UserID) on delete cascade
;

-- 
-- TABLE: SoftsResource 
--
ALTER TABLE SoftsResources ADD CONSTRAINT "RefSoftSkill10" 
    FOREIGN KEY (SoftSkillID)
    REFERENCES SoftSkills(SoftSkillID) on delete cascade
;

ALTER TABLE SoftsResources ADD CONSTRAINT "RefResourceData11" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade
;

-- 
-- TABLE: TechnologiesProject 
--
ALTER TABLE TechnologiesProjects ADD CONSTRAINT "RefProjects74" 
    FOREIGN KEY (ProjectID)
    REFERENCES Projects(ProjectID) on delete cascade
;

ALTER TABLE TechnologiesProjects ADD CONSTRAINT "RefHardSkill77" 
    FOREIGN KEY (HardSkillID)
    REFERENCES HardSkills(HardSkillID) on delete cascade
;

-- 
-- TABLE: User 
--
ALTER TABLE Users ADD CONSTRAINT "RefResourceData5" 
    FOREIGN KEY (ResourceDataID)
    REFERENCES ResourceDatas(ResourceDataID) on delete cascade;

ALTER TABLE Users ADD CONSTRAINT "RefRoles22" 
    FOREIGN KEY (RoleID)
    REFERENCES Roles(RoleID)
;


--
-- Check consult
--
Select TO_CHAR( daydate, 'Day'), daydate from days;
