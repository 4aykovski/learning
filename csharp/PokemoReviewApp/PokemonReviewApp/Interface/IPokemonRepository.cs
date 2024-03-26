using PokemoReviewApp.Models;

namespace PokemonReviewApp.Interface;

public interface IPokemonRepository
{
    ICollection<Pokemon> GetPokemons();

}