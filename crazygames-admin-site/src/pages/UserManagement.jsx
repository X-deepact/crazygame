import React, {useEffect, useState} from "react";
import { format, parseISO } from "date-fns";

import Constants from '@/components/Constants'
import DataTable from "@/components/DataTable";
import ConfirmDialog from "@/components/ConfirmDialog";
import InputDialog from "@/components/InputDialog";

// Helper functions to generate data
const randomDate = (start, end) => {
  return new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()))
    .toISOString()
    .split("T")[0];
};

const randomEmail = (username) => {
  const domains = ["example.com", "testmail.com", "mail.com", "domain.org"];
  return `${username.toLowerCase()}@${domains[Math.floor(Math.random() * domains.length)]}`;
};

const randomGender = () => {
  return ["M", "F", "O", "Z"][Math.floor(Math.random() * 4)];
};

// Generate sample data
const generateData = (numRecords) => {
  const startDate = new Date(2022, 0, 1);
  const endDate = new Date(2025, 0, 1);

  return Array.from({ length: numRecords }, (_, i) => ({
    ID: i + 1,
    Username: `User${i + 1}`,
    Email: randomEmail(`User${i + 1}`),
    Birthday: randomDate(startDate, endDate),
    Gender: randomGender(),
    UpdatedAt: new Date().toISOString().split("T")[0],
  }));
};

const UserManagement = () => {
  // Table Columns
  const columns = [
    { header: 'Username', field: 'Username', type: 'text', },
    { header: 'Email', field: 'Email', type: 'text', },
    { header: 'Birthday', field: 'Birthday', type: 'text', },
    { header: 'Gender', field: 'Gender', type: 'gender' },
    { header: 'Last Updated', field: 'UpdatedAt', type: 'text', },
    { header: 'Action', field: '', type: 'action' },
  ];

  const sampleData = generateData(1000);

  const [currentPage, setCurrentPage] = useState(1);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [totalPages, setTotalPages] = useState(1);
  const [tableData, setTableData] = useState([]);

  const handlePageChange = (pageNumber, rows, query) => {
    // TODO - integrate with API to load data by query, rows, page
    const total = Math.ceil(sampleData.length / rows);
    setTableData(sampleData.slice((pageNumber-1)*rows, pageNumber*rows));
    setTotalPages(total);
    setRowsPerPage(rows);
    setCurrentPage(pageNumber);
  };

  // input dialog
  const [isInputDialog, setIsInputDialog] = useState(false);
  const [currentData, setCurrentData] = useState({
      ID: '',
      Username: '',
      Email: '',
      Country: '',
      Birthday: '',
      Gender: '',
      Password: "",
      ConfirmPassword: "",
  });
  const [inputTitle, setInputTitle] = useState("");
  const [inputError, setInputError] = useState("");

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setCurrentData(prevState => ({
      ...prevState,
      [name]: value,
    }));
    setInputError("");
  };

  const validateForm = () => {
    if (currentData.Username==="") {
      setInputError("Please enter Username");
      return false;
    }

    if (currentData.Email==="") {
      setInputError("Please enter Email Address");
      return false;
    }

    if (currentData.Birthday==="") {
      setInputError("Please select your birthday");
      return false;
    }

    if (currentData.Birthday==="") {
      setInputError("Please select your birthday");
      return false;
    }

    const usernameRegex = /^[A-Za-z0-9._]{6,20}$/;
    if (!usernameRegex.test(currentData.Username)) {
      setInputError("Usernames must contain only letters, numbers, periods and underscores and have 6 - 20 characters.");
      return false;
    }

    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (!emailRegex.test(currentData.Email)) {
      setInputError("Please enter a valid email address.");
      return false;
    }

    return true;
  };

  const onAdd = () => {
    setInputError("");
    setCurrentData({
      ID: '',
      Username: '',
      Email: '',
      Country: '',
      Birthday: '',
      Gender: '',
      Password: "",
      ConfirmPassword: "",
    });
    setInputTitle("Add User");
    setIsInputDialog(true);
  };

  const onEdit = (record) => {
    setInputError("");
    setCurrentData(record);
    setInputTitle("Edit User");
    setIsInputDialog(true);
  };

  const onSave = () => {
    console.log("-------on save", currentData);
    if (!validateForm())
      return;

    // TODO - save user
    setIsInputDialog(false);
  };

  const onClose = () => {
    setIsInputDialog(false);
  };

  // delete dialog
  const [isDeleteConfirmDialog, setIsDeleteConfirmDialog] = useState(false);
  const onDelete = (record) => {
    setCurrentData(record);
    setIsDeleteConfirmDialog(true);
  };

  const onConfirmDelete = () => {
    setIsDeleteConfirmDialog(false);
    // TODO - integrate Delete API with currentData
  };

  // change password
  const [isPasswordDialog, setIsPasswordDialog] = useState(false);

  const onPassword = (record) => {
    setInputError("");
    setCurrentData({
      ID: record.ID,
      Username: '',
      Email: '',
      Country: '',
      Birthday: '',
      Gender: '',
      Password: "",
      ConfirmPassword: "",
    });
    setIsPasswordDialog(true);
  };

  const onSavePassword = () => {
    console.log("-------on save", currentData);

    if (currentData.Password == "") {
      setInputError("Please enter a Password");
      return false;
    }

    if (currentData.ConfirmPassword == "") {
      setInputError("Please enter a Confirm Password");
      return false;
    }

    if (currentData.ConfirmPassword != currentData.Password) {
      setInputError("Please enter valid password.");
      return false;
    }

    // TODO - reset password
    setIsPasswordDialog(false);
  };

  const onPaswordClose = () => {
    setIsPasswordDialog(false);
  };

  // initialize-------
  useEffect(() => {
    handlePageChange(1, 10, "");
  }, []);


  return (
    <div className="">
      <h1 className="text-2xl font-bold mb-4">User Management</h1>
      <p className="mb-6">Manage users here.</p>

      <DataTable
        columns={columns} data={tableData} currentPage={currentPage} rowsPerPage={rowsPerPage} totalPages={totalPages} onPageChange={handlePageChange}
        addLabel="Add User" onAdd={onAdd} onEdit={onEdit} onDelete={onDelete} onPassword={onPassword}
       />

      <ConfirmDialog title="Are you sure you want to delete this user?"
                     onOk={onConfirmDelete} onCancel={() => setIsDeleteConfirmDialog(false)}
                     isOpen={isDeleteConfirmDialog} onClose={() => setIsDeleteConfirmDialog(false)} />

      <InputDialog isOpen={isInputDialog} onClose={onClose} title={inputTitle} error={inputError} onOk={onSave} onCancel={onClose}>
        <div className="mb-4">
          <label htmlFor="Username" className="block text-sm font-semibold text-gray-700">Username</label>
          <input
            type="text"
            id="Username"
            name="Username"
            value={currentData?.Username || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="Email" className="block text-sm font-semibold text-gray-700">Username</label>
          <input
            type="email"
            id="Email"
            name="Email"
            value={currentData?.Email || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="Gender" className="block text-sm font-semibold text-gray-700">Gender</label>
          <select
            id="Gender"
            name="Gender"
            value={currentData?.Gender || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required>
            <option value="" disabled>Please select a gender</option>
            { Constants.GENDERS.map((col) => (
              <option value={col.value}>{col.name}</option>
            )) }
          </select>
        </div>

        <div className="mb-4">
          <label htmlFor="Country" className="block text-sm font-semibold text-gray-700">Country</label>
          <select
            id="Country"
            name="Country"
            value={currentData?.Country || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required>
            <option value="" disabled>Please select a country</option>
            { Constants.COUNTRIES.map((col) => (
              <option value={col}>{col}</option>
            )) }
          </select>
        </div>

        <div className="mb-4">
          <label htmlFor="Birthday" className="block text-sm font-semibold text-gray-700">Birthday</label>
          <input
            type="date"
            id="Birthday"
            name="Birthday"
            value={currentData?.Birthday || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>
      </InputDialog>


      <InputDialog isOpen={isPasswordDialog} onClose={onPaswordClose} title="Reset Password" onOk={onSavePassword} onCancel={onPaswordClose}>
        { inputError && (
          <h6 className="bg-red-100 text-red-500 rounded-md px-4 py-4 mb-5">{inputError}</h6>
        )}
        <div className="mb-4">
          <label htmlFor="Password" className="block text-sm font-semibold text-gray-700">Password</label>
          <input
            type="password"
            id="Password"
            name="Password"
            value={currentData?.Password || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="Email" className="block text-sm font-semibold text-gray-700">Confirm Password</label>
          <input
            type="password"
            id="ConfirmPassword"
            name="ConfirmPassword"
            value={currentData?.ConfirmPassword || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

      </InputDialog>
    </div>
  );
};

export default UserManagement;
