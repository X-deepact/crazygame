import React, {useState, useEffect} from "react";
import { useHistory } from "react-router-dom";
import Select from 'react-select';

const GameDetails = ({match}) => {
  const ID = match.params.id;
  const history = useHistory();

  const onBack = () => {
    history.push("/games");
  };

  const [activeTab, setActiveTab] = useState("main");
  const [isInputDialog, setIsInputDialog] = useState(false);

  const CategoryOptions = [
    { value: "action", label: "Action" },
    { value: "adventure", label: "Adventure" },
    { value: "rpg", label: "RPG" },
    { value: "alien", label: "Alien" },
    { value: "platform", label: "Platform" },
    { value: "hero", label: "Hero" },
    { value: "gun", label: "Gun" },
    { value: "survival", label: "Survival" }    
  ]

  const PlatformOptions = [
    { value: "action", label: "Action" },
    { value: "adventure", label: "Adventure" },
    { value: "rpg", label: "RPG" },
    { value: "strategy", label: "Strategy" },
    { value: "sports", label: "Sports" },    
  ]
  
  const [releasedData, setReleased] = useState({
    date: "",
    desc: ""
  });

  const [controlData, setControl] = useState({
    cont: ""
  });

  const [faqData, setFaq] = useState({
    title: "",
    answer: ""
  });

  const [formData, setFormData] = useState({
    main: {
      GameTitle: "",
      GameURL: "",
      ThumbnailURL:"",
      VideoURL:"",
      Description:"",
      HowToPlay:"",
      MoreGame:"",
      Category:[]
    },

    developer: {
      Name: "",
      Released: [],
      Technology:"",
      Platform:[],
      Controls:[],
      GameVideoURL:"",
    },

    faq: {
      Faq: [],
    },
  });

  const handleInputChange = (tab, field, value) => {
    console.log(tab, field, value)
    setFormData((prevData) => ({
      ...prevData,
      [tab]: {
        ...prevData[tab],
        [field]: value,
      },
    }));
    setInputError("");
  };

  const handleSave = () => {
    if(validateForm()) {
      console.log("Saved Data:", formData);

    } else {
      setIsInputDialog(false);
    }
   
  };


  // initialize-------
  useEffect(() => {
    // TODO - load game data
  }, [ID]);


  /** Released Function **/
  const addReleased = () => {
    let temp = {...formData};
    temp.developer.Released.push(releasedData);
    setFormData(temp);
    setReleased({
      date: "",
      desc: ""
    });
  }
  
  const editReleased = (index, field, value) => {
    const updatedReleased = formData.developer.Released.map((item, idx) =>
      idx === index ? { ...item, [field]: value } : item
    );
    setFormData({
      ...formData,
      developer: { ...formData.developer, Released: updatedReleased },
    });
  };

  const removeRleased = (idx) => {
    let temp = {...formData};
    temp.developer.Released.splice(idx, 1);
    setFormData(temp);
  }


  /** Controls Function **/
  const addControl = () => {
    let temp = {...formData};
    temp.developer.Controls.push(controlData);
    setFormData(temp);
    setControl({
        cont: ""
    });
  }
  
  const editControl = (index, field, value) => {
    const updatedControl = formData.developer.Controls.map((item, idx) =>
      idx === index ? { ...item, [field]: value } : item
    );
    setFormData({
      ...formData,
      developer: { ...formData.developer, Controls: updatedControl },
    });
  };

  const removeControl = (idx) => {
    let temp = {...formData};
    temp.developer.Controls.splice(idx, 1);
    setFormData(temp);
  }

  /** Faq Function **/
  const addFaq = () => {
    let temp = {...formData};
    temp.faq.Faq.push(faqData);
    setFormData(temp);
    setFaq({
        title: "",
        answer: ""
    });
  }

  const editFaq = (index, field, value) => {
    const updatedFaq = formData.faq.Faq.map((item, idx) =>
      idx === index ? { ...item, [field]: value } : item
    );
    setFormData({
      ...formData,
      faq: { ...formData.faq, Faq: updatedFaq },
    });
  };
  

  const removeFaq = (idx) => {
    let temp = {...formData};
    temp.faq.Faq.splice(idx, 1);
    setFormData(temp);
  }

  const [errorCard, setErrorCard] = useState(1);
  const [inputError, setInputError] = useState({
    main: "",
    developer: "",
    faq: ""
  });

  const validateForm = () => {
    
    let errFlag = 1;
    let errors = { main: "", developer: "", faq: "" };
    const { GameTitle, GameURL, ThumbnailURL, VideoURL, Description, Category, HowToPlay} = formData.main;
    const { Platform, GameVideoURL} = formData.developer;

    if (!GameTitle.trim()) {
      errors.main = "Please enter Game Title.";
    } else if (!GameURL.trim() || !/^(ftp|http|https):\/\/[^ "]+$/.test(GameURL)) {
      errors.main = !GameURL.trim() ? "Please enter Game URL." : "Please enter a valid Game URL.";
    } else if (!ThumbnailURL.trim() || !/^(ftp|http|https):\/\/[^ "]+$/.test(ThumbnailURL)) {
      errors.main = !ThumbnailURL.trim() ? "Please enter Thumbnail URL." : "Please enter a valid Thumbnail URL.";
    } else if (!VideoURL.trim() || !/^(ftp|http|https):\/\/[^ "]+$/.test(VideoURL)) {
      errors.main = !VideoURL.trim() ? "Please enter Video URL." : "Please enter a valid Video URL.";
    } else if (!Description.trim()) {
      errors.main = "Please enter Description.";
    } else if(!Category || Category.length === 0) {
      errors.main = "At least one Category must be selected."
    } else if(!HowToPlay.trim()) {
      errors.main = "Please enter HowToPlay."
      
      errFlag = 2;
    } else if(!Platform || Platform.length === 0) {
      errors.developer = "At least one platform must be selected."
    } else if (!GameVideoURL.trim() || !/^(ftp|http|https):\/\/[^ "]+$/.test(GameVideoURL)) {      
      errors.developer = !GameVideoURL.trim() ? "Please enter Game play Video URL." : "Please enter a valid Video URL.";      
    } 


    setInputError(errors);
    setErrorCard(errFlag);

    if ( errors.main != "")
      setActiveTab("main");
    else if (errors.developer != "")
      setActiveTab("developer");
    else if (errors.faq != "")
      setActiveTab("faq");

    return !Object.values(errors).some((error) => error);
  };

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4"><a className="cursor-pointer" onClick={onBack}><i className="fa fa-chevron-left"></i></a> Game Details</h1>
      <p className="mb-4">Manage game details.</p>

      {/* Tabs */}
      <div className="flex justify-between items-center pb-2 mb-4">
        <div className="flex space-x-4">
          <button
            className={`px-4 py-2 ${
              activeTab === "main"
                ? "text-teal-400 border-b-2 border-gray-400"
                : "text-gray-400"
            }`}
            onClick={() => setActiveTab("main")}
          >
            Main Information
          </button>

          <button
            className={`px-4 py-2 ${
              activeTab === "developer"
                ? "text-teal-400 border-b-2 border-gray-400"
                : "text-gray-400"
            }`}
            onClick={() => setActiveTab("developer")}
          >
            Developer Information
          </button>

          <button
            className={`px-4 py-2 ${
              activeTab === "faq"
                ? "text-teal-400 border-b-2 border-gray-400"
                : "text-gray-400"
            }`}
            onClick={() => setActiveTab("faq")}
          >
            FAQ
          </button>
        </div>

        {/* Save Button */}
        <button
          onClick={handleSave}
          className="bg-blue-600 hover:bg-blue-500 text-white px-4 py-2 rounded-lg"
        >
          Save
        </button>
      </div>

      {/* Tab Content */}
      <div>
        {activeTab === "main" && (
          <div className="flex flex-wrap gap-4">
            <div className="flex bg-white rounded-lg shadow-md p-6 w-full lg:w-[46%] md:w-[100%]">
              <div className="space-y-4 w-full">
                { errorCard == 1 && inputError.main && (
                  <div className="bg-red-100 text-red-500 rounded-md px-4 py-4 mb-5">
                    {inputError.main}
                  </div>
                )}
                <div>
                  <label className="block text-gray-600 mb-1">Game Title</label>
                  <input
                    type="text"
                    value={formData.main.GameTitle}
                    onChange={(e) =>
                      handleInputChange("main", "GameTitle", e.target.value)
                    }
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                  />                 
                </div>
                
                <div>
                  <label className="block text-gray-600 mb-1">Game URL</label>
                  <input
                    type="text"
                    value={formData.main.GameURL}
                    onChange={(e) =>
                      handleInputChange("main", "GameURL", e.target.value)
                    }
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                  />   
                </div>

                <div>
                  <label className="block text-gray-600 mb-1">Thumbnail URL</label>
                  <input
                    type="text"
                    value={formData.main.ThumbnailURL}
                    onChange={(e) =>
                      handleInputChange("main", "ThumbnailURL", e.target.value)
                    }
                    
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                  />     
                </div>

                <div>
                  <label className="block text-gray-600 mb-1">Video URL</label>
                  <input
                    type="text"
                    value={formData.main.VideoURL}
                    onChange={(e) =>
                      handleInputChange("main", "VideoURL", e.target.value)
                    }
                    
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                  />
                </div>
                
                <div>
                  <label className="block text-gray-600 mb-1">Description</label>
                  <textarea
                    value={formData.main.Description}
                    rows="5"
                    className="w-full px-4 py-2 rounded-lg border border-gray-400 focus:ring-2 focus:ring-teal-400 focus:outline-none resize-none"
                    placeholder="Write your message here..."
                    onChange={(e) =>
                      handleInputChange("main", "Description", e.target.value)
                    }                    
                  />                    
                </div>

                <div>
                  <label className="block text-gray-600 mb-1">Category</label>
                  <Select 
                    isMulti
                    name="Category"
                    options={CategoryOptions}
                    value={formData.main.Category }
                    onChange={(selectOptions) => 
                      handleInputChange("main", "Category", selectOptions)
                    }                    
                    className="w-full"
                  />
                  
                </div>
              </div>                
            </div>

            <div className="flex bg-white rounded-lg shadow-md p-6 w-full lg:w-[46%] md:w-[100%]">
              <div className="space-y- w-full">
                {errorCard == 2 && inputError.main && (
                  <div className="bg-red-100 text-red-500 rounded-md px-4 py-4 mb-5">
                    {inputError.main}
                  </div>
                )}
                <div>
                  <label className="block text-gray-600 mb-1">How to play</label>
                  <textarea
                    rows="15"
                    cols="100"
                    value={formData.main.HowToPlay}
                    className="w-full px-4 py-2 rounded-lg border border-gray-400 focus:ring-2 focus:ring-teal-400 focus:outline-none resize-none"
                    placeholder="Write your message here..."
                    onChange={(e) =>
                      handleInputChange("main", "HowToPlay", e.target.value)
                    }                    
                  />                  
                </div>

                <div className="mt-2">
                  <label className="block text-gray-600 mb-1">More Games</label>
                  <textarea
                    rows="5"
                    cols="100"
                    value={formData.main.MoreGame}
                    className="w-full px-4 py-2 rounded-lg border border-gray-400 focus:ring-2 focus:ring-teal-400 focus:outline-none resize-none"
                    placeholder="Write your message here..."
                    onChange={(e) =>
                      handleInputChange("main", "MoreGame", e.target.value)
                    }
                  />
                </div>
              </div>
            </div>

          </div>

        )}

        {activeTab === "developer" && (
           <div className="flex flex-wrap gap-4"> 
            <div className="bg-white rounded-lg shadow-md p-6 w-full lg:w-[46%] md:w-[100%]">
              <div className="space-y-4">
              {inputError.developer && (
                  <div className="bg-red-100 text-red-500 rounded-md px-4 py-4 mb-5">
                    {inputError.developer}
                  </div>
                )}
                <div>
                  <label className="block text-gray-600 mb-1">Name</label>
                  <input
                    type="text"
                    value={formData.developer.Name}
                    onChange={(e) =>
                        handleInputChange("developer", "Name", e.target.value)
                    }
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                  />
                </div>
               
                <div>
                  <label className="block text-gray-600 mb-1">Technology</label>
                  <input
                    type="text"
                    value={formData.developer.Technology}
                    onChange={(e) =>
                        handleInputChange("developer", "Technology", e.target.value)
                    }
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                  />
                </div>

                <div>
                  <label className="block text-gray-600 mb-1">Platform</label>
                  <Select 
                    isMulti
                    name="Platform"
                    options={PlatformOptions}
                    value={formData.developer.Platform }
                    onChange={(selectOptions) => 
                      handleInputChange("developer", "Platform", selectOptions)
                    }
                   
                    className="w-full"
                  />
                  
                </div>
                <div>
                  <label className="block text-gray-600 mb-1">Gameplay Video</label>
                  <input
                    type="text"
                    value={formData.developer.GameVideoURL}
                    onChange={(e) =>
                      handleInputChange("developer", "GameVideoURL", e.target.value)
                    }                    
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                    placeholder="Enter URL"
                  />                 
                </div>
              </div>                            
            </div>

            <div className="bg-white rounded-lg shadow-md p-6 w-full lg:w-[46%] md:w-[100%]">
              <div className="space-y-4">
                <div>
                  <label className="block text-gray-600 mb-1">Released Date</label>
                  <div className="flex gap-4">
                    <input
                      type="date"
                      id="date"
                      name="date"
                      value={releasedData.date}
                      onChange={(e)=> { setReleased({...releasedData, date: e.target.value}); }}
                      className="px-4 py-2 rounded-lg border border-gray-600 focus:outline-none focus:ring-2 focus:ring-teal-400"
                    />
                    <input
                      type="text"
                      value={releasedData.desc}
                      onChange={(e)=> { setReleased({...releasedData, desc: e.target.value}); }}
                      className="w-full px-4 py-2 rounded-lg border border-gray-400"
                    />
                    <button onClick={addReleased}>
                      <i className="fa fa-plus"></i>
                    </button>
                  </div>
                  {
                    formData.developer.Released.map((item, idx) => {
                      return <div className="flex gap-4 my-1">
                        <input
                          type="date"
                          id="date"
                          name="date"
                          value={item.date}
                          className="px-4 py-2 rounded-lg border border-gray-400 focus:outline-none focus:ring-2 focus:ring-teal-400"
                          onChange={(e) => editReleased(idx, 'date', e.target.value)}
                        />
                        <input
                          type="text"
                          value={item.desc}
                          className="w-full px-4 py-2 rounded-lg border border-gray-400"
                          onChange={(e) => editReleased(idx, 'desc', e.target.value)}             
                        />
                        <button onClick={() => removeRleased(idx)}>
                          <i className="fa fa-trash"></i>
                        </button>
                      </div>
                    })
                  }
                </div>
                
                <div>
                    <label className="block text-gray-600 mb-1">Controls</label>
                    <div className="flex gap-4">
                      <input
                        type="text"
                        value={controlData.cont}
                        onChange={(e)=> { setControl({...controlData, cont: e.target.value}); }}
                        className="w-full px-4 py-2 rounded-lg border border-gray-400"
                      />
                      <button onClick={addControl}>
                        <i className="fa fa-plus"></i>
                      </button>
                    </div>
                    {
                      formData.developer.Controls.map((item, idx) => {
                        return <div className="flex gap-4 my-1">
                          <input
                            type="text"
                            value={item.cont}
                            className="w-full px-4 py-2 rounded-lg border border-gray-400"
                            onChange={(e) => editControl(idx, 'cont', e.target.value)}             
                          />
                          <button onClick={ () => removeControl(idx)}>
                            <i className="fa fa-trash"></i>
                          </button>
                        </div>
                      })
                    }
                </div>                
              </div>
                            
            </div>
           </div>
        )}

        {activeTab === "faq" && (
          <div className="flex flex-wrap gap-4">
            <div className="bg-white rounded-lg shadow-md p-6 w-full lg:w-[46%] md:w-[100%]">
              <div className="flex gap-3">
                <div className="space-y-3 mr-2">
                  <div className="w-full">
                    <label className="block text-gray-600 mb-1">FAQ Title</label>
                    <input
                    type="text"
                    value={faqData.title}
                    onChange={(e) =>
                        { setFaq({...faqData, title:e.target.value}); }
                    }
                    className="w-full px-4 py-2 rounded-lg border border-gray-400"
                    />
                  </div>

                  <div className="w-full">
                    <label className="block text-gray-600 mb-1">Answer</label>
                    <textarea
                      rows="2"
                      cols="100"
                      value={faqData.answer}
                      className="w-full px-4 py-2 rounded-lg border border-gray-400 focus:ring-2 focus:ring-teal-400 focus:outline-none resize-none"
                      placeholder="Write your answer here..."
                      onChange={(e) =>
                        { setFaq({...faqData, answer:e.target.value}); }
                      }
                    />
                  </div>
                  
                </div>

                <div className="flex">
                  <button className="ml-auto" onClick={addFaq}>
                    <i className="fa fa-plus"></i>
                  </button>
                </div>
              </div>
                {
                  formData.faq.Faq.map((item, idx) => {
                    return <div className="my-3">
                      <div className="flex gap-4">
                        <div>
                          <input
                            type="text"
                            value={item.title}
                            className="w-full px-4 py-2 rounded-lg border border-gray-400"
                            onChange={(e) => editFaq(idx, 'title', e.target.value)}             
                          />
                          <textarea
                            type="text"
                            rows="2"
                            cols="100"
                            value={item.answer}
                            className="w-full px-4 py-2 my-1 rounded-lg border border-gray-400 focus:ring-2 focus:ring-teal-400 focus:outline-none resize-none"                                
                            onChange={(e) => editFaq(idx, 'answer', e.target.value)}             
                          />  
                        </div>

                        <button onClick={ () => removeFaq(idx)}>
                          <i className="fa fa-trash"></i>
                        </button>

                      </div>

                    </div>
                  })
                }
            </div>

            {/* <div className="bg-white rounded-lg shadow-md p-6 w-full lg:w-[46%] md:w-[100%]">
              <p>you can add other element here</p>
            </div> */}
          </div>
        )}
      </div>     
    </div>    
  );
};

export default GameDetails;
