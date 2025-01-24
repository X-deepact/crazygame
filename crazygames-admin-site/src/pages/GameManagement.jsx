import React, {useEffect, useState} from "react";
import { useHistory } from "react-router-dom";
import { format, parseISO } from "date-fns";

import Constants from '@/components/Constants'
import DataTable from "@/components/DataTable";
import ConfirmDialog from "@/components/ConfirmDialog";
import InputDialog from "@/components/InputDialog";

// Generate sample data
const generateData = (count) => {
  const games = [];
  for (let i = 1; i <= count; i++) {
    games.push({
      ID: i,
      GameTitle: `Game Title ${i}`,
      GameURL: `https://example.com/game-${i}`,
      Developer: `Developer ${i}`,
      ThumbnailURL: `https://imgs.crazygames.com/farm-merge-valley_16x9/20250110081325/farm-merge-valley_16x9-cover?auto=format%2Ccompress&q=90&cs=strip&w=273&fit=crop`,
      HoverVideoUrl: `https://www.w3schools.com/html/mov_bbb.mp4`,
      ReleaseDate: `202${Math.floor(i % 10)}/${
        (i % 12) + 1
      }/${String(i % 28 + 1).padStart(2, '0')}`,
      UpdatedAt: `202${Math.floor(i % 10)}/${
        ((i + 2) % 12) + 1
      }/${String((i + 3) % 28 + 1).padStart(2, '0')}`,
    });
  }
  return games;
};

const GameManagement = () => {
  const history = useHistory();

  // Table Columns
  const columns = [
    { header: 'Title', field: 'GameTitle', type: 'text', },
    { header: 'Game URL', field: 'GameURL', type: 'text' },
    { header: 'Developer', field: 'Developer', type: 'text', },
    { header: 'Thumbnail', field: 'ThumbnailURL', type: 'image' },
//    { header: 'Video', field: 'HoverVideoUrl', type: 'video' },
    { header: 'Release Date', field: 'ReleaseDate', type: 'text' },
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
    if (currentData.GameTitle === "") {
      setInputError("Please enter Title");
      return false;
    }

    if (currentData.Description === "") {
      setInputError("Please enter Description");
      return false;
    }

    if (currentData.Developer === "") {
      setInputError("Please enter Developer");
      return false;
    }

    const regex = /^(ftp|http|https):\/\/[^ "]+$/;
    if (!regex.test(currentData.ThumbnailURL)) {
      setInputError("Please enter a valid url for Thumbnail");
      return false;
    }

    if (!regex.test(currentData.HoverVideoUrl)) {
      setInputError("Please enter a valid url for Hover Video");
      return false;
    }

    if (!regex.test(currentData.GameURL)) {
      setInputError("Please enter a valid url for Game");
      return false;
    }

    return true;
  };

  const onAdd = () => {
    setInputError("");
    setCurrentData({
      ID: '',
      GameTitle: '',
      Description: '',
      Developer: '',
      ThumbnailURL: '',
      HoverVideoUrl: '',
      GameURL: "",
    });
    setInputTitle("Add Game");
    setIsInputDialog(true);
  };

  const onEdit = (record) => {
    console.log("--------", record);
    history.push(`/games/${record.ID}`); // Navigate to the game page with the ID
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

  // initialize-------
  useEffect(() => {
    handlePageChange(1, 10, "");
  }, []);


  return (
    <div className="">
      <h1 className="text-2xl font-bold mb-4">Game Management</h1>
      <p className="mb-6">Manage games here.</p>

      <DataTable
        columns={columns} data={tableData} currentPage={currentPage} rowsPerPage={rowsPerPage} totalPages={totalPages} onPageChange={handlePageChange}
        addLabel="Add Game" onAdd={onAdd} onEdit={onEdit} onDelete={onDelete}
       />

      <ConfirmDialog title="Are you sure you want to delete this game?"
                     onOk={onConfirmDelete} onCancel={() => setIsDeleteConfirmDialog(false)}
                     isOpen={isDeleteConfirmDialog} onClose={() => setIsDeleteConfirmDialog(false)} />

      <InputDialog isOpen={isInputDialog} onClose={onClose} title={inputTitle} error={inputError} onOk={onSave} onCancel={onClose}>
        <div className="mb-4">
          <label htmlFor="GameTitle" className="block text-sm font-semibold text-gray-700">Title</label>
          <input
            type="text"
            id="GameTitle"
            name="GameTitle"
            value={currentData?.GameTitle || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="Description" className="block text-sm font-semibold text-gray-700">Description</label>
          <textarea
            id="Description"
            name="Description"
            value={currentData?.Description || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="Developer" className="block text-sm font-semibold text-gray-700">Developer</label>
          <input
            type="text"
            id="Developer"
            name="Developer"
            value={currentData?.Developer || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="ThumbnailURL" className="block text-sm font-semibold text-gray-700">Thumbnail URL</label>
          <input
            type="text"
            id="ThumbnailURL"
            name="ThumbnailURL"
            value={currentData?.ThumbnailURL || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="HoverVideoUrl" className="block text-sm font-semibold text-gray-700">Hover Video URL</label>
          <input
            type="text"
            id="HoverVideoUrl"
            name="HoverVideoUrl"
            value={currentData?.HoverVideoUrl || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="GameURL" className="block text-sm font-semibold text-gray-700">Game URL</label>
          <input
            type="text"
            id="GameURL"
            name="GameURL"
            value={currentData?.GameURL || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>

      </InputDialog>

    </div>
  );
};

export default GameManagement;
