#ifndef _XG_DRIVERAPI_H_ 
#define _XG_DRIVERAPI_H_ 

#ifdef WIN32 
#define XG_API __cdecl
 
#else
//linuxs about
#define XG_API
#endif
typedef long long int64;
 
#ifdef __cplusplus
extern "C"{
#endif 

#define  XG_C_NULL                    0
#define  XG_C_BOOL                    1
#define  XG_C_CHAR                    2
#define  XG_C_TINYINT                 3
#define  XG_C_SHORT                   4
#define  XG_C_INTEGER                 5
#define  XG_C_BIGINT                  6
#define  XG_C_FLOAT                   7
#define  XG_C_DOUBLE                  8
#define  XG_C_NUMERIC                 9
#define  XG_C_DATE	                  10
#define  XG_C_TIME			          11
#define  XG_C_TIME_TZ                 12
#define  XG_C_DATETIME                13
#define  XG_C_DATETIME_TZ             14
#define  XG_C_BINARY                  15

#define  XG_C_NCHAR                   18 

#define  XG_C_INTERVAL                21
#define  XG_C_INTERVAL_YEAR_TO_MONTH  28
#define  XG_C_INTERVAL_DAY_TO_SECOND  31

#define  XG_C_TIMESTAMP     XG_C_DATETIME
#define  XG_C_LOB                     40
#define  XG_C_CLOB                    41
#define  XG_C_BLOB                    42

#define XG_SUCCESS              0
#define XG_NO_DATA              100  

#define XG_ERROR               -1
#define XG_NET_ERROR           -4
#define XG_INVALID_ARG         -3
#define XG_SOCKET_ERROR        -8 
#define XG_LOGIN_ERROR         -9 

#define XG_NULL_DATA           -11
#define XG_TRUNCATED_DATA      -12
#define XG_DATATYPE_ERROR      -13      /* Data type cannot be converted */  
#define XG_FLOW_DATA           -14      /* Data type out of bounds */
#define XG_COL_SEQ_ERR         -15      /* Data serial number out of bounds */
#define XG_COL_EXCEPT_DATAOFF  -18      /* Data offset out of bounds */ 

#define XG_COL_DATA_OVERFLOW    98   

typedef enum tagPARAMINOUT_TYPE{
    PARAM_INPUT=1,
    PARAM_OUTPUT=2,
    PARAM_INPUTOUTPUT=3,
    PARAM_RETURNVALUE=6,

}PARAMINOUT_TYPE;
 
/*  
 * CONNECT
 *
 *  */
int XG_API XGC_OpenConn(char* Conn_str,void** p_conn);
int XG_API XGC_CloseConn(void** p_conn);

/* 
 * IPS CONNECT
 * */
int XG_API XGC_OpenConn_Ips(char* Conn_str,int ntimes,void** turnIP_attrs,void** p_conn);


/*
 *  PARAMETER
 *
 * */

/* Explicitly create parameters */
int XG_API XGC_CreateParams(void** p_params);

/* Reset parameters in connection 
 *
 * Note: explicitly created parameters will not be processed
 *
 * */
int XG_API XGC_ResetParams(void** p_conn);

/* Bind explicitly created parameters to the connection */
int XG_API XGC_BindParams2Conn(void** p_conn,void** p_params);

int XG_API XGC_BindParamByName(void** p_conn, char* param_name, int param_type, 
           int datatype, void*  value, int param_size,  int* rt_code,  int* rlen_val);

int XG_API XGC_BindParamArrayByName(void** p_conn, char* param_name,int param_num, 
     int param_type,int datatype, int array_size, void* array_value, 
           int param_size, int * rlen_val);

int XG_API XGC_BindParamByPos(void** p_conn, int param_no,int param_type, 
                int datatype, void* value, int param_size, int * rlen_val);

int XG_API XGC_BindParamArrayByPos(void** p_conn, int param_no, int param_num, 
     int param_type,int datatype, int array_size, void* array_value, int param_size, int * rlen_val); 
 

/*
 * COMMAND
 *
 * */
int XG_API XGC_Execute_no_query(void** p_conn,char* cmd_sql); 

int XG_API XGC_Execute_query_with_one(void** p_conn ,char* cmd_sql,void* re_val,int* type);

/* SELECT WITH SERVER CURSOR */
int XG_API XGC_Prepare2(void** p_conn,char* cmd_sql,char* prepare_name); 
int XG_API XGC_Execute2(void** p_conn, char* prepare_name, char* servercursor_name,void** pres);
int XG_API XGC_ExecBatch(void**  p_conn,char* cmd_sql, int ArrayCount);
int XG_API XGC_UnPrepare(void** p_conn,char* prepare_name);
int XG_API XGC_CloseCursor(void** p_conn,char* cursor_name);
int XG_API XGC_ExecwithDataReader(void** p_conn ,char* cmd_sql,void** p_res,int* field_num,int64* rowcount,int* effected_num);

/* 
 * Get result set from server cursor
 * int XG_API XGC_FetchServerCursorRowset(void** p_conn ,char* cmd_sql,void** p_res);
 *
 * */
int XG_API XGC_FetchServerCursorRowset(void** p_conn ,char* servercursor_name,void** p_res);

/* Fetching data from the server cursor header (extra)*/
int XG_API XGC_FetchRefCursorHead(void** p_conn ,char* Cursor_name ,void** p_res,int* field_num,int64* rowcount,int* cached);

/* XGC_Prepare2() + XGC_Execute2() (extra) */
int XG_API XGC_ExecwithServerCursorReader(void** p_conn ,char* cmd_sql, 
    char* Cursor_name ,void** p_res,int* field_num,int64* rowcount,int* effected_num);

/*
 * Execution of stored procedures and functions, 
 * involving input and output of parameters
 * */
int XG_API XGC_Execute_procesure(void** p_conn , char*  cmd_sql,void* para); 

/*
 *  RESULT
 *
 * */
int XG_API  XGC_GetData(void** pTr_Result,int col_no,int TarCtype, 
           void* TarValuePtr,int BuffLen,int* lenPtr);

int XG_API  XGC_getResultcolType(void**  pTr_Result,int col_no,int* col_type) ;
int XG_API  XGC_getResultcolname(void**  pTr_Result,int col_no,char* col_name) ;
int XG_API  XGC_getResultcolseq(void**  pTr_Result,char* col_name);
int XG_API  XGC_getResultColumnsnum(void**  pTr_Result,int* field_num);
int XG_API  XGC_getResultRecordnum(void**  pTr_Result,int* record_num);
int XG_API  XGC_getResultcolmodi(void**  pTr_Result, int col_no, int* modi);//add 202-02-19

int XG_API  XGC_getResultColInfo(void**  pTr_Result,int col_no, 
    char* col_Tabname, char* col_name, char* col_alias, int* datatype,
         int* col_modi,int* col_flag);

/* Result set cursor moves to the next result set */
int XG_API XGC_ReadNext(void** p_res);

/* Release result set */
int XG_API XGC_FreeRowset(void** p_res);

/*
 * Get the next result set, suitable for multiple result sets
 * */
int XG_API XGC_NextResult(void** p_res);

/* Attribute */
int XG_API XGC_GetAttr(void** hd_ptr, int attrtype, void * ValuePtr, int  BuffLen, int* ret_attr_type, int* re_len);
int XG_API XGC_SetAttr(void** hd_ptr, int attrtype, const void * ValuePtr, int  BuffLen);

/*
 * BLOB\CLOB
 * */
int XG_API XGC_Create_Lob(void** Lob_ptr);
int XG_API XGC_Put_Lob_data(void** Lob_ptr, void* data, int len );
int XG_API XGC_Get_Lob_data(void** Lob_ptr, void* data, int len);
int XG_API XGC_Distroy_Lob(void** Lob_ptr);
int XG_API XGC_LobWrite_SetPos(void** Lob_ptr,int setpos);
int XG_API XGC_LobRead_SetPos(void** Lob_ptr,int setpos);
int XG_API XGC_Reset_Lob(void** Lob_ptr);

/* 
 * ERROR INFO
 * */
int  XG_API XGC_GetError(void** hd_ptr, char* err_text,int* rlen);
int  XG_API XGC_GetErrorInfo(void** p_handptr, char* ccode, char* errmessage, int* rlen); 
int  XG_API XGC_GetErrorInfoOption(void** p_handptr, char* ccode, int * ret_code, char* errmessage, int max_message_len, int* rlen);

/*
 *  OTHER
 *
 * */
void XG_API XGC_FreePtr(void**Ptr);
void XG_API XGC_Drop(void**Ptr_obj);
int  XG_API dt2dtm_Api(long long  t,char * p_dt);
int  XG_API Release_IpsAttrs(void** pconn_IpsAttr);//ips= 
int  XG_API fun_sql_type(char* sql);
int  XG_API XGC_Reset(void**Ptr_obj); 

int  XG_API  XGC_getResultRet(void**  pTr_Result,int * type, 
       int* field_num,int * rowcount, int *effected_num ,char* insert_rowid);

/* Get the rowid of the last insert operation */
int  XG_API XGC_GetLastInsertId(void** p_conn, char* insert_rowid);

#ifdef __cplusplus

}
#endif 

#endif
