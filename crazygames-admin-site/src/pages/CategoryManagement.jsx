import React, {useEffect, useState} from "react";

import DataTable from "@/components/DataTable";
import ConfirmDialog from "@/components/ConfirmDialog";
import InputDialog from "@/components/InputDialog";

const generateData = (count) => {
  const categoryTitles = [
    'Action', 'Adventure', 'Basketball', 'Beauty', 'Bike', 'Car', 'Card', 'Casual',
    'Clicker', 'Controller', 'Dress Up', 'Driving', 'Escape', 'Flash', 'FPS', 'Horror',
    '.io', 'Mahjong', 'Minecraft', 'Pool', 'Puzzle', 'Shooting', 'Soccer', 'Sports', 'Stickman'
  ];

  const descriptions = [
    'The latest in action-packed adventures.',
    'Explore thrilling adventures and stories.',
    'Basketball games and challenges.',
    'Beauty and fashion tips.',
    'Bike racing and challenges.',
    'Drive and race cars in thrilling environments.',
    'Play card games with friends.',
    'Casual games for relaxation.',
    'Clicker games for simple fun.',
    'Controller-based games for fun challenges.',
    'Dress up games with various outfits.',
    'Drive vehicles and complete missions.',
    'Escape rooms and puzzles to solve.',
    'Flash games for short bursts of fun.',
    'FPS games with intense action and shooting.',
    'Horror games for those who dare.',
    '.io games that you can play with others.',
    'Mahjong puzzle games.',
    'Minecraft-style sandbox games.',
    'Pool and billiards games.',
    'Puzzle-solving games to test your brain.',
    'Shooting games with various weapons and missions.',
    'Soccer games with realistic gameplay.',
    'Sports games for all enthusiasts.',
    'Stickman-themed action games.'
  ];

  const categories = [];
  for (let i = 1; i <= count; i++) {
    const randomCategoryIndex = Math.floor(Math.random() * categoryTitles.length);
    const randomTitle = `${categoryTitles[randomCategoryIndex]} ${i}`;
    const randomDescription = descriptions[Math.floor(Math.random() * descriptions.length)];
    const updatedAt = new Date(2022, Math.floor(Math.random() * 12), Math.floor(Math.random() * 28) + 1).toLocaleDateString();

    // Generate an icon URL from placeimg.com
    const icon = `https://imgs.crazygames.com/icon/${categoryTitles[randomCategoryIndex]}.svg`;  // Category-based index for variety

    categories.push({
      ID: i,
      CategoryName: randomTitle,
      Icon: icon,
      Description: randomDescription,
      UpdatedAt: updatedAt,
    });
  }
  return categories;
};

const CategoryManagement = () => {
  // Table Columns
  const columns = [
    { header: 'Category Name', field: 'CategoryName', type: 'text', },
    { header: 'Icon', field: 'Icon', type: 'image', },
    { header: 'Description', field: 'Description', type: 'text', },
    { header: 'Last Updated', field: 'UpdatedAt', type: 'text', },
    { header: 'Action', field: '', type: 'action', },
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
    CategoryTitle: '',
    Description: '',
    Icon: ''
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
    if (currentData.CategoryTitle==="") {
      setInputError("Please enter Category Title");
      return false;
    }

    if (currentData.Icon==="") {
      setInputError("Please enter Category Icon");
      return false;
    }

    const regex = /^(ftp|http|https):\/\/[^ "]+$/;
    if (!regex.test(currentData.Icon)) {
      setInputError("Please enter a valid url for Icon");
      return false;
    }

    return true;
  };

  const onAdd = () => {
    setInputError("");
    setCurrentData({
      ID: '',
      CategoryTitle: '',
      Description: '',
      Icon: ''
    });
    setInputTitle("Add Category");
    setIsInputDialog(true);
  };

  const onEdit = (record) => {
    setInputError("");
    setCurrentData(record);
    setInputTitle("Edit Category");
    setIsInputDialog(true);
  };

  const onSave = () => {
    console.log("-------on save", currentData);
    if (!validateForm())
      return;

    // TODO - save category
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
      <h1 className="text-2xl font-bold mb-4">Category Management</h1>
      <p className="mb-6">Manage categories here.</p>

      <DataTable
        columns={columns} data={tableData} currentPage={currentPage} rowsPerPage={rowsPerPage} totalPages={totalPages} onPageChange={handlePageChange}
        addLabel="Add Category" onAdd={onAdd} onEdit={onEdit} onDelete={onDelete}
       />

      <ConfirmDialog title="Are you sure you want to delete this category?"
                     onOk={onConfirmDelete} onCancel={() => setIsDeleteConfirmDialog(false)}
                     isOpen={isDeleteConfirmDialog} onClose={() => setIsDeleteConfirmDialog(false)} />

      <InputDialog isOpen={isInputDialog} onClose={onClose} title={inputTitle} error={inputError} onOk={onSave} onCancel={onClose}>
        <div className="mb-4">
          <label htmlFor="CategoryName" className="block text-sm font-semibold text-gray-700">Category Title</label>
          <input
            type="text"
            id="CategoryName"
            name="CategoryName"
            value={currentData?.CategoryName || ''}
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
          <label htmlFor="Icon" className="block text-sm font-semibold text-gray-700">Icon URL</label>
          <input
            type="text"
            id="Icon"
            name="Icon"
            value={currentData?.Icon || ''}
            onChange={handleInputChange}
            className="w-full p-2 border border-gray-300 rounded"
            required
          />
        </div>
      </InputDialog>

    </div>
  );
};

export default CategoryManagement;
